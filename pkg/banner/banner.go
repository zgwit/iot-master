package banner

import "fmt"

const banner = `
###        #######       #     #                                   
 #            #          ##   ##   ##    ####  ##### ###### #####  
 #    ####    #          # # # #  #  #  #        #   #      #    # 
 #   #    #   #   #####  #  #  # #    #  ####    #   #####  #    # 
 #   #    #   #          #     # ######      #   #   #      #####  
 #    ####    #          #     # #    # #    #   #   #      #   #  
###           #          #     # #    #  ####    #   ###### #    # 

`

//Print 打印 banner 和 附加信息
func Print(args ...string) {
	fmt.Print(banner)
	for _, s := range args {
		fmt.Println(s)
	}
}
