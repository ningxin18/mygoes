package main

import (
	"encoding/hex"
	"fmt"
	"encoding/pem"
	"os"
	"io/ioutil"
)

func main(){

	rca:="2d2d2d2d2d424547494e2043455254494649434154452d2d2d2d2d0a4d49494457444343416b436741774942416749514f5346533550635342756544776c4c546f657a646644414e42676b71686b69473977304241517346414441670a4d517377435159445651514745774a44546a45524d4138474131554543684d4956584e6c51326868615734774868634e4d5467774e7a45774d5441774d7a51300a5768634e4d6a67774e7a45774d5441774d7a5130576a41674d517377435159445651514745774a44546a45524d4138474131554543684d4956584e6c513268680a61573477676745694d4130474353714753496233445145424151554141344942447741776767454b416f4942415144426865306754664a4c6b646639637479790a4b716f635476304e6f664973526d713067796e355a616c6832734470306d38465073356243663067357063485168373250436e65794a714c3875364f736b7a750a4d5548433374724458652b425130534d77553576313772416d745a5a396c545747635652386661584c5a744762305137434634626b4f4a58703465547475394f0a4461457948386b71664556576155776e5a754b4738395a6b4933766761346135784a78724c5a5644764e705069547a69316c653370496a7552426f506b2b4f590a625a72627441463838754d6c47635978594344686e7070384c6c546c315568496b346364716a586e744c694b5342304f75756872445373664e715a32304467530a36424e4171586e42667431744a4b56434e473645766d30633077766d316a64456530444143645245482b4f326f765a4a412b704c797a2b464a6a6a506e5876690a44576d5641674d424141476a6759307767596f7744675944565230504151482f4241514441674b6b4d423047413155644a5151574d42514743437347415155460a42774d42426767724267454642516344416a415042674e5648524d4241663845425441444151482f4d45674741315564455152424d442b4343577876593246730a6147397a644945794d4867794d5745334d7a4931515463314e446b79524749785a5755344e6b4d785a444a454d6a49354f4452694f444e6d4e6a41344d6d55310a514445324d79356a623230774451594a4b6f5a496876634e4151454c42514144676745424141385067544a2f62397556333046756b78544250594d3466314b500a787335746d362b4754654c2f414b302b664e66555a314938546e3754724650726f597934524662337a325437466f58414e33597943656257307452344f354e790a424d7531544b386f4d484152654852636e4e4764425353614b2b6749324d594779336a79426975385a47696d2f78662b776441696d2b5466474e59736c7476720a4b77434b5441514c7051542b78723270662f5a384d616e584474764e414e483632474557784835464f673665466e43513469357169704770396c576e4d4262770a44694f79494c36314e654a655852326443694c496d474f6c534d76415a6a492b5035456854633347312b4452532f7343777a656d733974486d4a4c59356c4c390a3776424e3864734c42776b77783461456538613854693646387779487142665076646b6261454f636d65595734676f5a7a5237694f6f455a594a593d0a2d2d2d2d2d454e442043455254494649434154452d2d2d2d2d0a"
	certToByte,err:=hex.DecodeString(rca)
	if err !=nil {
		fmt.Println("user's certificate format error")
	}

	Block, _:= pem.Decode(certToByte)
	if Block == nil {
		fmt.Println("ecaFile error")
	}

	SaveCA("asdfasdfasdfd",Block.Bytes)
	cer,_:=getCert()
	fmt.Println(cer)
}


func SaveCA(caName string,cert []byte ){
	certOut, err := os.Create(caName)
	if err != nil {
		fmt.Println("failed to open for writing: %s", err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert})
	fmt.Println("write ecert to", certOut.Name())
	certOut.Close()
}

func getCert() (string,error) {

	cert := "asdfasdfasdfd"
	//解析rca证书
	certByte, err := ioutil.ReadFile(cert)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return "",err
	}
	certAscii:=hex.EncodeToString(certByte[:])
	return certAscii,err
}
