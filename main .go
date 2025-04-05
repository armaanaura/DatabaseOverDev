package main
import(
	"encoding/binary"
	"fmt"
	"os"
)
func main(){
	filename :="devDb.data"
	file, err:=os.Create(filename)
	if(err!=nil){
		fmt.Println("Error while creating the file", err)
		return
	}
	defer file.Close()
	page:=make([]byte, 16*1024)
	fmt.Println("Page size:", len(page))
	binary.LittleEndian.PutUint32(page[0:4],1)
	binary.LittleEndian.PutUint16(page[4:6],0)
	page[6]=1

	_,err= file.Write(page)
	if(err!=nil){
		fmt.Println("Error while writing the file", err)
		return
	}
	fmt.Println("File written successfully")

}
