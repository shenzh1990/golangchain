package main

import (
	"github.com/gotoeasy/glang/cmn"
	"github.com/tmc/langchaingo/llms/ollama"
	"golangchain/service"
)

var chatService service.ChatService

func main() {
	llm, err := ollama.New(ollama.WithModel("llama2-chinese:13b"), ollama.WithServerURL("http://192.168.100.12:11434"))
	if err != nil {
		cmn.Error(err)
	}
	//chatService.AddDoc(llm, "/Users/shenzehua/Desktop/1.pdf")
	chatService.ChatWithDoc(llm, "你是谁？")
	//cmn.SetGlcClient(cmn.NewGlcClient(&cmn.GlcOptions{Enable: false, LogLevel: "INFO"})) // 控制台INFO日志级别输出
	//runtime.GOMAXPROCS(settings.CpuMaxProcess)
	//// 显式调用触发数据库、redis等
	////redisutil.Start()
	//onstart.Run()
	//llm, err := ollama.New(ollama.WithModel("llama2-chinese:13b"), ollama.WithServerURL("http://192.168.100.12:11434"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//ctx := context.Background()
	//content := []llms.MessageContent{
	//	llms.TextParts(llms.ChatMessageTypeSystem, "你是一个医生助手，并且只会使用中文回答。"),
	//	llms.TextParts(llms.ChatMessageTypeHuman, "我现在有点头疼，我应该怎么办"),
	//}
	//completion, err := llm.GenerateContent(ctx, content, llms.WithTemperature(1), llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
	//	fmt.Print(string(chunk))
	//	return nil
	//}))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_ = completion
}
