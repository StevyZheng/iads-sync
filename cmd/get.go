package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"iads/lib/linux"
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getCpuInfoCmd)
	getCmd.AddCommand(getMbInfoCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get info",
}

var getCpuInfoCmd = &cobra.Command{
	Use:   "cpuinfo",
	Short: "Print the cpu info",
	Run: func(cmd *cobra.Command, args []string) {
		cpuinfo := new(linux.CpuHwInfo)
		cpuinfo.GetCpuHwInfo()
		fmt.Println("model:", cpuinfo.Model)
		fmt.Println("sockets:", cpuinfo.Count)
		fmt.Println("cores:", cpuinfo.CoreCount)
		fmt.Println("stepping:", cpuinfo.Stepping)
	},
}

var getMbInfoCmd = &cobra.Command{
	Use:   "mbinfo",
	Short: "Print the motherborad info",
	Run: func(cmd *cobra.Command, args []string) {
		mbinfo := new(linux.MotherboradInfo)
		mbinfo.GetMbInfo()
		fmt.Println("model:", mbinfo.Model)
		fmt.Println("biosVer:", mbinfo.BiosVer)
		fmt.Println("biosDate:", mbinfo.BiosDate)
	},
}
