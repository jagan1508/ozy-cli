/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/jagan1508/ozy/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// modifyCmd represents the modify command
var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Used to modify a Todo item",
	Long:  `This command helps you to modify an item in the todo list`,
	Run:   modifyRun,
}

var (
	doneSt  string
	doneUsr string
	priSt   int
	textSt  string
)

func doneP(items []todo.Item, i int) string {
	if items[i-1].Done {
		doneUsr = "done"
	} else {
		doneUsr = "not done"
	}
	return doneUsr
}

func modifyRun(cmd *cobra.Command, args []string) {
	items, _ := todo.ReadItems(viper.GetString("datafile"))
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], " is not a valid label \n", err)
	}

	if i > 0 && i <= len(items) {
		fmt.Println("Modyfying ", items[i-1].Label()+"\t"+items[i-1].PrettyDone()+"\t"+items[i-1].PrettyP()+"\t"+items[i-1].Text+"\t")
		fmt.Println("Press Enter if you dont want to change a field ")

		fmt.Println("The original status is ", doneP(items, i))
		fmt.Print("Enter ( change ) to modify the status: ")
		fmt.Scanln(&doneSt)
		if doneSt == "change" {

			items[i-1].Done = !items[i-1].Done
			fmt.Println("Modified status to ", doneP(items, i))
		}

		fmt.Println("The original item priority is ", items[i-1].Priority)
		fmt.Print("Enter new priority: ")
		fmt.Scanln(&priSt)
		if priSt != items[i-1].Priority && priSt >= 1 && priSt <= 3 {
			items[i-1].Priority = priSt
			fmt.Println("Modified priority to ", priSt)
		} else {
			fmt.Println("Invalid priority")
		}

		fmt.Println("The original item content is ", items[i-1].Text)
		fmt.Print("Enter new content: ")
		inputReader := bufio.NewReader(os.Stdin)
		textSt, _ := inputReader.ReadString('\n')
		textSt = strings.TrimSpace(textSt)
		if textSt != items[i-1].Text && textSt != "" {
			items[i-1].Text = textSt
			fmt.Println("Modified todo content to ", textSt)
		} else {
			fmt.Println("You didnt change anything !!")
		}
		sort.Sort(todo.ByPri(items))
		todo.SaveItems(viper.GetString("datafile"), items)
	}
}

func init() {
	rootCmd.AddCommand(modifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
