package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nikola43/web3golanghelper/web3helper"
	pancakeRouter "boopaibot/contracts/IPancakeRouter02"
)

const (
	// BotToken is the Telegram Bot API token.
	//BotToken = "5960863904:AAFdz0O5CwzglyFJ_Hz_nTkJqiZ-zUuBZ_8"
	BotToken = "5829617229:AAHHqca1eu2PG3FS-bXATVEptV2Y9ky4H-M"
	// ChatID is the Telegram chat ID.
	//ChatID = -891755767
	// Call buybot command
	CallBuybot = "/buybot"
	// Call buybot config command
	CallBuybotConfig = "/buybotconfig"
)

func main() {

	web3GolangHelper := initWeb3()
	//migrate(db)
	routerAddress := "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
	routerAbi, _ := abi.JSON(strings.NewReader(string(pancakeFactory.PancakeABI)))

	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			// get the first argument of the command
			// and use it as the text of the reply
			msg.Text = "Hi " + update.Message.CommandArguments()

			//msg.Text = update.Message.
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

func initWeb3() *web3helper.Web3GolangHelper {
	pk := "b366406bc0b4883b9b4b3b41117d6c62839174b7d21ec32a5ad0cc76cb3496bd"
	rpcUrl := "https://speedy-nodes-nyc.moralis.io/84a2745d907034e6d388f8d6/bsc/testnet"
	wsUrl := "wss://speedy-nodes-nyc.moralis.io/84a2745d907034e6d388f8d6/bsc/testnet/ws"
	web3GolangHelper := web3helper.NewWeb3GolangHelper(rpcUrl, wsUrl, pk)

	chainID, err := web3GolangHelper.HttpClient().NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Chain Id: " + chainID.String())
	return web3GolangHelper
}
