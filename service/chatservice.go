package service

import (
	"context"
	"github.com/gotoeasy/glang/cmn"
	"github.com/jackc/pgx/v5"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
	"golangchain/common"
	"golangchain/pkg/settings"
)

type ChatService struct{}

func (cs *ChatService) AddDoc(llm *ollama.LLM, dirFile string) {
	e, err := embeddings.NewEmbedder(llm)
	if err != nil {
		cmn.Error(err)
	}
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, settings.AppConfig.Db.DBUrl)
	store, err := pgvector.New(
		ctx,
		pgvector.WithConn(conn),
		pgvector.WithEmbedder(e),
	)
	if err != nil {
		cmn.Error(err)
	}
	docs, err := common.FileToChunks(dirFile, 768, 64)
	if err != nil {
		cmn.Error(err)
	}
	_, err = store.AddDocuments(context.Background(), docs)
	if err != nil {
		cmn.Error(err)
	}
}
func (cs *ChatService) ChatWithDoc(llm *ollama.LLM, prompt string) {
	e, err := embeddings.NewEmbedder(llm)
	if err != nil {
		cmn.Error(err)
	}
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, settings.AppConfig.Db.DBUrl)
	store, err := pgvector.New(
		ctx,
		pgvector.WithConn(conn),
		pgvector.WithEmbedder(e),
	)
	if err != nil {
		cmn.Error(err)
	}

	// Search for similar documents.vectorstores.WithScoreThreshold(0.80)
	docs_s, err := store.SimilaritySearch(ctx, prompt, 5)
	if err != nil {
		cmn.Error(err)
	}
	// 创建一个新的聊天消息历史记录
	history := memory.NewChatMessageHistory()
	// 将检索到的文档添加到历史记录中
	for _, doc := range docs_s {
		history.AddAIMessage(ctx, doc.PageContent)
	}
	// 使用历史记录创建一个新的对话缓冲区
	conversation := memory.NewConversationBuffer(memory.WithChatHistory(history))
	//executor := agents.NewExecutor(
	//	agents.NewConversationalAgent(llm, nil),
	//	nil,
	//	agents.WithMemory(conversation),
	//)
	// 设置链调用选项
	options := []chains.ChainCallOption{
		chains.WithTemperature(0.8),
	}
	translatePrompt := prompts.NewPromptTemplate(
		"请你从已知信息{{.data}}中找出下面问题的答案，问题是 {{.text}} ,如果找不到问题答案，你可以提示未在指定内容终获取您要的答案。",
		[]string{"data", "text"},
	)
	llmChain := chains.NewLLMChain(llm, translatePrompt)

	// Otherwise the call function must be used.
	outputValues, err := chains.Call(ctx, llmChain, map[string]any{
		"data": conversation.ChatHistory,
		"text": prompt,
	}, options...)
	// 运行链
	//res, err := chains.Run(ctx, llmChain, prompt, options...)
	if err != nil {
		cmn.Error(err)
	}
	cmn.Info(outputValues)
}
