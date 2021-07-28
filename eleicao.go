package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Message struct {
	Body      string
	Timestamp [5]float64
	Distancia float64
	Emissor int
	pidMenor int
	LeaderPid int
}

type Matrix [][]float64

func event(pid int, counter*[5]float64) {
    counter[pid-1] += 1
	if (pid-1)==0{
		fmt.Printf("%v",*counter)
		fmt.Printf("Evento interno do processo P.\n")
	}else{
		if (pid-1)==1{
			fmt.Printf("%v",*counter)
			fmt.Printf("Evento interno do processo Q.\n")
		}else{
			if (pid-1)==2{
				fmt.Printf("%v",*counter)
				fmt.Printf("Evento interno do processo S.\n")
			}else{
				if (pid-1)==3{
					fmt.Printf("%v",*counter)
					fmt.Printf("Evento interno do processo T.\n")
				}else{
					if (pid-1)==4{
						fmt.Printf("%v",*counter)
						fmt.Printf("Evento interno do processo R.\n")
					}
				}
			}
		}
	}
}

func max(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

func calcTimestamp(recvTimestamp, counter float64) float64 {
	return max(recvTimestamp, counter)
}

func sendMessage(ch []chan Message, pid int, counter*[5]float64, channelIndex int) {
	counter[pid-1]+=1

if (pid-1)==0{
		counter[0]+=1
			
		if(channelIndex==0){
			ch[channelIndex] <- Message{Body:"ENVIADA PELO PROCESSO P", Timestamp:*counter}
			fmt.Printf("%v",*counter)
			fmt.Printf("Mensagem enviada do PROCESSO P para O PROCESSO Q.\n")
			
		}
		if(channelIndex==2){/*2ps*/
			ch[channelIndex] <- Message{Body:"ENVIADA PELO PROCESSO P", Timestamp:*counter}
			fmt.Printf("%v",*counter)
			fmt.Printf("Mensagem enviada do PROCESSO P para O PROCESSO S.\n")
		}
		
	}else{
		if (pid-1)==1{
			counter[1]+=1
				
			if(channelIndex==1){/*0qp*/
				ch[channelIndex] <- Message{Body:"ENVIADA PELO PROCESSO Q", Timestamp:*counter}
				fmt.Printf("%v",*counter)
				fmt.Printf("Mensagem enviada do PROCESSO P para O PROCESSO P.\n")
			}
			if(channelIndex==4){/*2qr*/
				ch[channelIndex] <- Message{Body:"ENVIADA PELO PROCESSO Q", Timestamp:*counter}
				fmt.Printf("%v",*counter)
				fmt.Printf("Mensagem enviada do PROCESSO P para O PROCESSO T.\n")
			}
			
		}else{
			if (pid-1)==2{
				counter[2]+=1
					
				if(channelIndex==3){/*3sp*/
					ch[channelIndex] <- Message{Body:"ENVIADA PELO PROCESSO S", Timestamp:*counter}
					fmt.Printf("%v",*counter)
					fmt.Printf("Mensagem enviada do PROCESSO P para O PROCESSO P.\n")
				}
				if(channelIndex==5){/*10st*/
					ch[channelIndex] <- Message{Body:"ENVIADA PELO PROCESSO S", Timestamp:*counter}
					fmt.Printf("%v",*counter)
					fmt.Printf("Mensagem enviada do PROCESSO P para O PROCESSO T.\n")
				}
					if(channelIndex==6){/*10st*/
					ch[channelIndex] <- Message{Body:"ENVIADA PELO PROCESSO S", Timestamp:*counter}
					fmt.Printf("%v",*counter)
					fmt.Printf("Mensagem enviada do PROCESSO P para O PROCESSO T.\n")
				}
					if(channelIndex==8){/*10st*/
					ch[channelIndex] <- Message{Body:"ENVIADA PELO PROCESSO S", Timestamp:*counter}
					fmt.Printf("%v",*counter)
					fmt.Printf("Mensagem enviada do PROCESSO P para O PROCESSO T.\n")
				}
			}else{
				if (pid-1)==3{
					counter[2]+=1
						
					if(channelIndex==7){/*5tp*/
						ch[channelIndex] <- Message{Body:"ENVIADA PELO PROCESSO T", Timestamp:*counter}
						fmt.Printf("%v",*counter)
						fmt.Printf("Mensagem enviada do PROCESSO P para O PROCESSO P.\n")
					}
					if(channelIndex==10){/*7tq*/
						ch[channelIndex] <- Message{Body:"ENVIADA PELO PROCESSO T", Timestamp:*counter}
						fmt.Printf("%v",*counter)
						fmt.Printf("Mensagem enviada do PROCESSO P para O PROCESSO Q.\n")
					}
					
				}else{
					if (pid-1)==4{
						counter[2]+=1
							
						if(channelIndex==9){/*9rq*/
							ch[channelIndex] <- Message{Body:"ENVIADA PELO PROCESSO R", Timestamp:*counter}
							fmt.Printf("%v",*counter)
							fmt.Printf("Mensagem enviada do PROCESSO Q para O PROCESSO Q.\n")
						}
						if(channelIndex==11){/*9rq*/
							ch[channelIndex] <- Message{Body:"ENVIADA PELO PROCESSO R", Timestamp:*counter}
							fmt.Printf("%v",*counter)
							fmt.Printf("Mensagem enviada do PROCESSO Q para O PROCESSO Q.\n")
						}
					}
				}
			}
		}
	}
}

func disparochandyMisra(ch []chan Message, pid int) {
	ch[12+pid-1] <- Message{Body:"ROUTER", Emissor:pid}
}

func inicializadorTarry(ch []chan Message, pid int,lider*int) {
	ch[12+29+pid-1] <- Message{Body:"TOKEN", Emissor:pid, LeaderPid:*lider,pidMenor:0xf3f3f3f3}
}

func receiveMessage(ch []chan Message, pid int, counter*[5]float64, channelReceptor int) {
	message := <-ch[channelReceptor]
	counter[0] = calcTimestamp(counter[0], message.Timestamp[0]);
	counter[1] = calcTimestamp(counter[1], message.Timestamp[1]);
	counter[2] = calcTimestamp(counter[2], message.Timestamp[2]);
	counter[3] = calcTimestamp(counter[3], message.Timestamp[3]);
	counter[4] = calcTimestamp(counter[4], message.Timestamp[4]);
	
	counter[pid-1]+=1

	if (pid-1)==0{
		counter[0]+=1
			
		fmt.Printf("%v",*counter)
		fmt.Printf("MENSAGEM RECEBIDA NO PROCESSO P %s.\n", message.Body)
	}else{
		if (pid-1)==1{
			counter[1]+=1
				
			fmt.Printf("%v",*counter)
			fmt.Printf("MENSAGEM RECEBIDA NO PROCESSO Q %s.\n", message.Body)
		}else{
			if (pid-1)==2{
				counter[2]+=1
				
				fmt.Printf("%v",*counter)
				fmt.Printf("MENSAGEM RECEBIDA NO PROCESSO S %s.\n", message.Body)
			}else{
				if (pid-1)==3{
					counter[3]+=1
					
					fmt.Printf("%v",*counter)
					fmt.Printf("MENSAGEM RECEBIDA NO PROCESSO T %s.\n", message.Body)
				}else{
					if (pid-1)==4{
						counter[4]+=1
						
						fmt.Printf("%v",*counter)
						fmt.Printf("MENSAGEM RECEBIDA NO PROCESSO R %s.\n", message.Body)
					}
				}
			}
		}
	}
}

func mapCanais(emissor int, receptor int) int{
	var concatenaNumero int
	concatenaNumero=10*emissor+receptor
	switch {
		case concatenaNumero==12:
			return 0
		case concatenaNumero==21:
			return 1
		case concatenaNumero==13:
			return 2
		case concatenaNumero==31:
			return 3
		case concatenaNumero==23:
			return 4
		case concatenaNumero==32:
			return 5
		case concatenaNumero==34:
			return 6
		case concatenaNumero==43:
			return 7
		case concatenaNumero==35:
			return 8
		case concatenaNumero==53:
			return 9
		case concatenaNumero==45:
			return 10
		case concatenaNumero==54:
			return 11
		default:
			return -1
	}
}

func callbackTarry(ch []chan Message, vizinhosTarryP*[]int, vizinhosTarryQ*[]int, vizinhosTarryR*[]int, vizinhosTarryS*[]int, vizinhosTarryT*[]int, primeiroTokenRecebidoP*bool, primeiroTokenRecebidoQ*bool, primeiroTokenRecebidoR*bool, primeiroTokenRecebidoS*bool, primeiroTokenRecebidoT*bool){
	for{
		message := <-ch[46]
		if(message.Body=="callback"){
			(*vizinhosTarryP)[0]=-1
			(*vizinhosTarryP)[1]=0
			(*vizinhosTarryP)[2]=0
			(*vizinhosTarryP)[3]=-1
			(*vizinhosTarryP)[4]=-1
			*primeiroTokenRecebidoP=false

			(*vizinhosTarryQ)[0]=0
			(*vizinhosTarryQ)[1]=-1
			(*vizinhosTarryQ)[2]=0
			(*vizinhosTarryQ)[3]=-1
			(*vizinhosTarryQ)[4]=-1
			*primeiroTokenRecebidoQ=false

			(*vizinhosTarryR)[0]=0
			(*vizinhosTarryR)[1]=0
			(*vizinhosTarryR)[2]=-1
			(*vizinhosTarryR)[3]=0
			(*vizinhosTarryR)[4]=0
			*primeiroTokenRecebidoR=false

			(*vizinhosTarryS)[0]=-1
			(*vizinhosTarryS)[1]=-1
			(*vizinhosTarryS)[2]=0
			(*vizinhosTarryS)[3]=-1
			(*vizinhosTarryS)[4]=0
			*primeiroTokenRecebidoS=false

			(*vizinhosTarryT)[0]=-1
			(*vizinhosTarryT)[1]=-1
			(*vizinhosTarryT)[2]=0
			(*vizinhosTarryT)[3]=0
			(*vizinhosTarryT)[4]=-1
			*primeiroTokenRecebidoT=false
		}
	}
}

func TarryGerenciador(ch []chan Message, pid int, counter*[5]float64, channelReceptor int, primeiroTokenRecebido*bool, parent*int, neighbors*[]int, lider*int,eleicaoEmAndamento*bool) {
	var indiceSorteado int
	mapProcessos := []string{"P","Q","R","S","T"}
	var contVizinhos int
    for{
		message := <-ch[channelReceptor]
		if (message.Body=="TOKEN"){
			if(message.Emissor==pid&&!*primeiroTokenRecebido){
				*primeiroTokenRecebido=true
				fmt.Printf("	[Tarry]%vINICIADO NO PROCESSO %s\n", *counter, mapProcessos[pid-1])
			}else{
				fmt.Printf("	[Tarry]%vO PROCESSO %s RECEBEU TOKEN DO PROCESSO %s\n", *counter, mapProcessos[pid-1], mapProcessos[message.Emissor-1])
			}
			if(message.Emissor!=pid&&!*primeiroTokenRecebido){
				*primeiroTokenRecebido=true
				*parent=message.Emissor
				*lider=message.LeaderPid
				fmt.Printf("	[Tarry]%vO ESTADO NO PROCESSO %s é: %v E O parent É: %s\n", *counter, mapProcessos[pid-1], *counter, mapProcessos[*parent-1])
			}
			contVizinhos=0
			for j := 0; j < 5; j++ {
				if((*neighbors)[j]!=-1){
					contVizinhos=contVizinhos+1
				}
			}
			if(contVizinhos>0){
				rand.Seed(time.Now().UnixNano())
				min := 1
				max := contVizinhos
				indiceSorteado=rand.Intn(max - min + 1) + min
			}else{
				contVizinhos=0
			}
			if(contVizinhos!=0){
				var j int
				j=0
				for j < 5 {
					if((*neighbors)[j]!=-1){
						indiceSorteado=indiceSorteado-1
						if(indiceSorteado==0){
							if(j+1!=message.Emissor){
							    if(pid<message.pidMenor){
							        ch[mapCanais(pid, j+1)] <- Message{Body:"TOKEN", Emissor:pid, LeaderPid:*lider,pidMenor:pid}
							    }else{
								ch[mapCanais(pid, j+1)] <- Message{Body:"TOKEN", Emissor:pid,LeaderPid:*lider,pidMenor:message.pidMenor}
							    }
								(*neighbors)[j]=-1
							}else{
								contVizinhos=0
								for j := 0; j < 5; j++ {
									if((*neighbors)[j]!=-1){
										contVizinhos=contVizinhos+1
									}
								}
								if(j+1==message.Emissor&&contVizinhos==1&&(*neighbors)[j]!=-1){
								    if(pid<message.pidMenor){
								    	ch[mapCanais(pid, j+1)] <- Message{Body:"TOKEN", Emissor:pid,LeaderPid:*lider,pidMenor:pid}
								    }else{
								   
					                    	ch[mapCanais(pid, j+1)] <- Message{Body:"TOKEN", Emissor:pid,LeaderPid:*lider,pidMenor:message.pidMenor}
								    }
									(*neighbors)[j]=-1
								} else{
									contVizinhos=0
									for j := 0; j < 5; j++ {
										if((*neighbors)[j]!=-1){
											contVizinhos=contVizinhos+1
										}
									}
									rand.Seed(time.Now().UnixNano())
									min := 1
									max := contVizinhos
									indiceSorteado=rand.Intn(max - min + 1) + min
									j=-1
								}
							}
						}
					}
					j=j+1
				}
			}else{
				
				if(*eleicaoEmAndamento){
				    *eleicaoEmAndamento=false
				    *lider=message.pidMenor
					fmt.Printf("	[Tarry]%vELEIÇÃO DE LÍDER FINALIZADA NO PROCESSO %s O LÍDER FOI %s\n", *counter, mapProcessos[pid-1],mapProcessos[*lider-1])
					ch[46] <- Message{Body:"callback"}
					time.Sleep(1 * time.Second)
					inicializadorTarry(ch, pid,lider)
					
					disparochandyMisra(ch,*lider)
				}
			}
		}
	}
}

func chandyMisraGerenciador(ch []chan Message, pid int, counter*[5]float64, channelReceptor int, distancia*float64, parent*int) {
	neighbors:= Matrix{{-1,0.5,0.5,-1,-1},
	                        {0.5,-1,0.5,-1,-1},
	                        {0.5,0.5,-1,0.125,0.125},
	                        {-1,-1,0.125,-1,0.125},
	                        {-1,-1,0.125,0.125,-1}}
	                        /*pp, pq, pr, ps, pt
	                          qp, qq, qr, qs, qt
	                          rp, rq, rr, rs, rt
	                          sp, sq, sr, ss, st
	                          tp, tq, tr, ts, tt*/
	mapProcessos := []string{"P","Q","R","S","T"}
	for{
		message := <-ch[channelReceptor]
		if (message.Body=="ROUTER"){
			if(message.Emissor==pid){
				*distancia=0
				fmt.Printf("	[chandyMisra]%vINICIADO NO PROCESSO %s: dist %f\n", *counter, mapProcessos[pid-1], *distancia)
				for j := 0; j < 5; j++ {
					if(neighbors[pid-1][j]!=-1){
						ch[mapCanais(pid, j+1)] <- Message{Body:"ROUTER", Emissor:pid, Distancia:*distancia}
					}
				}
			}else{
				if(message.Distancia+neighbors[pid-1][message.Emissor-1]<*distancia){
					*parent=message.Emissor
					*distancia=message.Distancia+neighbors[pid-1][message.Emissor-1]
					fmt.Printf("	[chandyMisra]%vO PROCESSO %s RECEBEU ROUTER DO PROCESSO %s: dist %f parent %s\n", *counter, mapProcessos[pid-1], mapProcessos[message.Emissor-1], *distancia, mapProcessos[*parent-1])
					for j := 0; j < 5; j++ {
						if(neighbors[pid-1][j]!=-1&&j!=*parent-1){
							ch[mapCanais(pid, j+1)] <- Message{Body:"ROUTER", Emissor:pid, Distancia:*distancia}
						}
					}
				}
			}
		}else{
			if(message.Body=="TOKEN"){
				ch[channelReceptor+29]<-message
			}else{
				ch[channelReceptor+17]<-message
				receiveMessage(ch, pid, counter, channelReceptor+17)
			}
		}
	}
}

func processP(ch []chan Message, vizinhosTarryP*[]int, primeiroTokenRecebidoP*bool) {
	pid := 1
	counter := [5]float64{0,0,0,0,0}
	var distancia float64
	var parent int
	var parentTravessia int
	var lider int
	var eleicaoEmAndamento bool
	parentTravessia=-1
	distancia=0xf3f3f3f3
	parent=-1
	mapProcessos := []string{"P","Q","R","S","T"}
	lider=-1
	eleicaoEmAndamento=true
	go chandyMisraGerenciador(ch, pid, &counter, 1, &distancia, &parent)
	go chandyMisraGerenciador(ch, pid, &counter, 3, &distancia, &parent)
	go chandyMisraGerenciador(ch, pid, &counter, 12, &distancia, &parent)
	go TarryGerenciador(ch, pid, &counter, 1+29, primeiroTokenRecebidoP, &parentTravessia, vizinhosTarryP, &lider,&eleicaoEmAndamento)
	go TarryGerenciador(ch, pid, &counter, 3+29, primeiroTokenRecebidoP, &parentTravessia, vizinhosTarryP, &lider,&eleicaoEmAndamento)
	go TarryGerenciador(ch, pid, &counter, 12+29, primeiroTokenRecebidoP, &parentTravessia, vizinhosTarryP, &lider,&eleicaoEmAndamento)
	//disparochandyMisra(ch, pid)
	
	event(pid, &counter)
	sendMessage(ch, pid, &counter, 0)
	event(pid, &counter)
	sendMessage(ch, pid, &counter, 2)
	time.Sleep(4* time.Second)
	fmt.Printf("FINAL %s %v lider: %s\n", mapProcessos[pid-1], counter, mapProcessos[lider-1])
	if(parent!=-1){
		fmt.Printf("[chandyMisra]FINAL %s %v dist:%f parent:%s\n", mapProcessos[pid-1], counter, distancia, mapProcessos[parent-1])
		if(parentTravessia!=-1){
			fmt.Printf("[Tarry]FINAL %s %v parent:%s\n", mapProcessos[pid-1], counter, mapProcessos[parentTravessia-1])
		}
	}else{
		fmt.Printf("[chandyMisra]FINAL %s %v dist:%f\n", mapProcessos[pid-1], counter, distancia)
	}
}

func processQ(ch []chan Message, vizinhosTarryQ*[]int, primeiroTokenRecebidoQ*bool) {
	pid := 2
	counter := [5]float64{0,0,0,0,0}
	var distancia float64
	var parent int
	var parentTravessia int
	var lider int
	var eleicaoEmAndamento bool
	parentTravessia=-1
	distancia=0xf3f3f3f3
	parent=-1
	mapProcessos := []string{"P","Q","R","S","T"}
		lider=-1
	eleicaoEmAndamento=true
	go chandyMisraGerenciador(ch, pid, &counter, 0, &distancia, &parent)
	go chandyMisraGerenciador(ch, pid, &counter, 5, &distancia, &parent)
	go chandyMisraGerenciador(ch, pid, &counter, 13, &distancia, &parent)
	go TarryGerenciador(ch, pid, &counter, 0+29, primeiroTokenRecebidoQ, &parentTravessia, vizinhosTarryQ, &lider,&eleicaoEmAndamento)
	go TarryGerenciador(ch, pid, &counter, 5+29, primeiroTokenRecebidoQ, &parentTravessia, vizinhosTarryQ, &lider,&eleicaoEmAndamento)
	go TarryGerenciador(ch, pid, &counter, 13+29, primeiroTokenRecebidoQ, &parentTravessia, vizinhosTarryQ, &lider,&eleicaoEmAndamento)
	event(pid, &counter)
	sendMessage(ch, pid, &counter, 1)
	event(pid, &counter)
	sendMessage(ch, pid, &counter, 4)
	time.Sleep(4 * time.Second)
	fmt.Printf("FINAL %s %v lider: %s\n", mapProcessos[pid-1], counter, mapProcessos[lider-1])
	if(parent!=-1){
		fmt.Printf("[chandyMisra]FINAL %s %v dist:%f parent:%s\n", mapProcessos[pid-1], counter, distancia, mapProcessos[parent-1])
		if(parentTravessia!=-1){
			fmt.Printf("[Tarry]FINAL %s %v parent:%s\n", mapProcessos[pid-1], counter, mapProcessos[parentTravessia-1])
		}
	}else{
		fmt.Printf("[chandyMisra]FINAL %s %v dist:%f\n", mapProcessos[pid-1], counter, distancia)
	}
}

func processR(ch []chan Message, vizinhosTarryR*[]int, primeiroTokenRecebidoR*bool) {
	pid := 3
	counter := [5]float64{0,0,0,0,0}
	var distancia float64
	var parent int
	var parentTravessia int
	var lider int
	var eleicaoEmAndamento bool
	parentTravessia=-1
	distancia=0xf3f3f3f3
	parent=-1
	mapProcessos := []string{"P","Q","R","S","T"}
	lider=-1
	eleicaoEmAndamento=true
	go chandyMisraGerenciador(ch, pid, &counter, 2, &distancia, &parent)
	go chandyMisraGerenciador(ch, pid, &counter, 4, &distancia, &parent)
	go chandyMisraGerenciador(ch, pid, &counter, 7, &distancia, &parent)
	go chandyMisraGerenciador(ch, pid, &counter, 9, &distancia, &parent)
	go chandyMisraGerenciador(ch, pid, &counter, 14, &distancia, &parent)
	go TarryGerenciador(ch, pid, &counter, 2+29, primeiroTokenRecebidoR, &parentTravessia, vizinhosTarryR, &lider,&eleicaoEmAndamento)
	go TarryGerenciador(ch, pid, &counter, 4+29, primeiroTokenRecebidoR, &parentTravessia, vizinhosTarryR, &lider,&eleicaoEmAndamento)
	go TarryGerenciador(ch, pid, &counter, 7+29, primeiroTokenRecebidoR, &parentTravessia, vizinhosTarryR, &lider,&eleicaoEmAndamento)
	go TarryGerenciador(ch, pid, &counter, 9+29, primeiroTokenRecebidoR, &parentTravessia, vizinhosTarryR, &lider,&eleicaoEmAndamento)
	go TarryGerenciador(ch, pid, &counter, 14+29, primeiroTokenRecebidoR, &parentTravessia, vizinhosTarryR, &lider,&eleicaoEmAndamento)
	event(pid, &counter)
	sendMessage(ch, pid, &counter, 3)
	event(pid, &counter)
	sendMessage(ch, pid, &counter, 5)
	event(pid, &counter)
	sendMessage(ch, pid, &counter, 6)
	event(pid, &counter)
	sendMessage(ch, pid, &counter, 8)
	time.Sleep(4 * time.Second)
	fmt.Printf("FINAL %s %v lider: %s\n", mapProcessos[pid-1], counter, mapProcessos[lider-1])
	if(parent!=-1){
		fmt.Printf("[chandyMisra]FINAL %s %v dist:%f parent:%s\n", mapProcessos[pid-1], counter, distancia, mapProcessos[parent-1])
		if(parentTravessia!=-1){
			fmt.Printf("[Tarry]FINAL %s %v parent:%s\n", mapProcessos[pid-1], counter, mapProcessos[parentTravessia-1])
		}
	}else{
		fmt.Printf("[chandyMisra]FINAL %s %v dist:%f\n", mapProcessos[pid-1], counter, distancia)
	}
}

func processS(ch []chan Message, vizinhosTarryS*[]int, primeiroTokenRecebidoS*bool) {
	pid := 4
	counter := [5]float64{0,0,0,0,0}
	var distancia float64
	var parent int
	var parentTravessia int
	var lider int
	var eleicaoEmAndamento bool
	parentTravessia=-1
	distancia=0xf3f3f3f3
	parent=-1
	mapProcessos := []string{"P","Q","R","S","T"}
	lider=-1
	eleicaoEmAndamento=true
	go chandyMisraGerenciador(ch, pid, &counter, 6, &distancia, &parent)
	go chandyMisraGerenciador(ch, pid, &counter, 11, &distancia, &parent)
	go chandyMisraGerenciador(ch, pid, &counter, 15, &distancia, &parent)
	go TarryGerenciador(ch, pid, &counter, 6+29, primeiroTokenRecebidoS, &parentTravessia, vizinhosTarryS, &lider,&eleicaoEmAndamento)
	go TarryGerenciador(ch, pid, &counter, 11+29, primeiroTokenRecebidoS, &parentTravessia, vizinhosTarryS, &lider,&eleicaoEmAndamento)
	go TarryGerenciador(ch, pid, &counter, 15+29, primeiroTokenRecebidoS, &parentTravessia, vizinhosTarryS, &lider,&eleicaoEmAndamento)
	event(pid, &counter)
	sendMessage(ch, pid, &counter, 7)
	event(pid, &counter)
	sendMessage(ch, pid, &counter, 10)
	time.Sleep(4* time.Second)
	fmt.Printf("FINAL %s %v lider: %s\n", mapProcessos[pid-1], counter, mapProcessos[lider-1])
	if(parent!=-1){
		fmt.Printf("[chandyMisra]FINAL %s %v dist:%f parent:%s\n", mapProcessos[pid-1], counter, distancia, mapProcessos[parent-1])
		if(parentTravessia!=-1){
			fmt.Printf("[Tarry]FINAL %s %v parent:%s\n", mapProcessos[pid-1], counter, mapProcessos[parentTravessia-1])
		}
	}else{
		fmt.Printf("[chandyMisra]FINAL %s %v dist:%f\n", mapProcessos[pid-1], counter, distancia)
	}
}

func processT(ch []chan Message, vizinhosTarryT*[]int, primeiroTokenRecebidoT*bool) {
	pid := 5
	counter := [5]float64{0,0,0,0,0}
	var distancia float64
	var parent int
	var parentTravessia int
	var lider int
	var eleicaoEmAndamento bool
	parentTravessia=-1
	distancia=0xf3f3f3f3
	parent=-1
	mapProcessos := []string{"P","Q","R","S","T"}
	lider=-1
	eleicaoEmAndamento=true
	go chandyMisraGerenciador(ch, pid, &counter, 8, &distancia, &parent)
	go chandyMisraGerenciador(ch, pid, &counter, 10, &distancia, &parent)
	go chandyMisraGerenciador(ch, pid, &counter, 16, &distancia, &parent)
	go TarryGerenciador(ch, pid, &counter, 8+29, primeiroTokenRecebidoT, &parentTravessia, vizinhosTarryT, &lider,&eleicaoEmAndamento)
	go TarryGerenciador(ch, pid, &counter, 10+29, primeiroTokenRecebidoT, &parentTravessia, vizinhosTarryT, &lider,&eleicaoEmAndamento)
	go TarryGerenciador(ch, pid, &counter, 16+29, primeiroTokenRecebidoT, &parentTravessia, vizinhosTarryT, &lider,&eleicaoEmAndamento)
	inicializadorTarry(ch, pid,&lider)
	event(pid, &counter)
	sendMessage(ch, pid, &counter, 9)
	event(pid, &counter)
	sendMessage(ch, pid, &counter,11)
	time.Sleep(4 * time.Second)
	fmt.Printf("FINAL %s %v lider: %s\n", mapProcessos[pid-1], counter, mapProcessos[lider-1])
	if(parent!=-1){
		fmt.Printf("[chandyMisra]FINAL %s %v dist:%f parent:%s\n", mapProcessos[pid-1], counter, distancia, mapProcessos[parent-1])
		if(parentTravessia!=-1){
			fmt.Printf("[Tarry]FINAL %s %v parent:%s\n", mapProcessos[pid-1], counter, mapProcessos[parentTravessia-1])
		}
	}else{
		fmt.Printf("[chandyMisra]FINAL %s %v dist:%f\n", mapProcessos[pid-1], counter, distancia)
	}
}

func main() {
	var chans = []chan Message {
        make(chan Message, 100), /*chandyMisra0pq*/
		make(chan Message, 100), /*chandyMisra1qp*/
        make(chan Message, 100), /*chandyMisra2pr*/
        make(chan Message, 100), /*chandyMisra3rp*/
		make(chan Message, 100), /*chandyMisra4qr*/
		make(chan Message, 100), /*chandyMisra5rq*/
		make(chan Message, 100), /*chandyMisra6rs*/
		make(chan Message, 100), /*chandyMisra7sr*/
		make(chan Message, 100), /*chandyMisra8rt*/
		make(chan Message, 100), /*chandyMisra9tr*/
		make(chan Message, 100), /*chandyMisra10st*/
		make(chan Message, 100), /*chandyMisra11ts*/
        make(chan Message, 100), /*chandyMisra12p*/
        make(chan Message, 100), /*chandyMisra13q*/
		make(chan Message, 100), /*chandyMisra14r*/
        make(chan Message, 100), /*chandyMisra15s*/
		make(chan Message, 100), /*chandyMisra16t*/
		make(chan Message, 100), /*RECIEVE17pq*/
		make(chan Message, 100), /*RECIEVE18qp*/
        make(chan Message, 100), /*RECIEVE19pr*/
        make(chan Message, 100), /*RECIEVE20rp*/
		make(chan Message, 100), /*RECIEVE21qr*/
		make(chan Message, 100), /*RECIEVE22rq*/
		make(chan Message, 100), /*RECIEVE23rs*/
		make(chan Message, 100), /*RECIEVE24sr*/
		make(chan Message, 100), /*RECIEVE25rt*/
		make(chan Message, 100), /*RECIEVE26tr*/
		make(chan Message, 100), /*RECIEVE27st*/
		make(chan Message, 100), /*RECIEVE28ts*/
		make(chan Message, 100), /*TRAVESSIA29pq*/
		make(chan Message, 100), /*TRAVESSIA30qp*/
        make(chan Message, 100), /*TRAVESSIA31pr*/
        make(chan Message, 100), /*TRAVESSIA32rp*/
		make(chan Message, 100), /*TRAVESSIA33qr*/
		make(chan Message, 100), /*TRAVESSIA34rq*/
		make(chan Message, 100), /*TRAVESSIA35rs*/
		make(chan Message, 100), /*TRAVESSIA36sr*/
		make(chan Message, 100), /*TRAVESSIA37rt*/
		make(chan Message, 100), /*TRAVESSIA38tr*/
		make(chan Message, 100), /*TRAVESSIA39st*/
		make(chan Message, 100), /*TRAVESSIA40ts*/
        make(chan Message, 100), /*TRAVESSIA41p*/
        make(chan Message, 100), /*TRAVESSIA42q*/
		make(chan Message, 100), /*TRAVESSIA43r*/
        make(chan Message, 100), /*TRAVESSIA44s*/
		make(chan Message, 100), /*TRAVESSIA45t*/
		make(chan Message, 100), /*callbackTarry46t*/
    }
	vizinhosTarryP:= []int{-1,0,0,-1,-1}
	vizinhosTarryQ:= []int{0,-1,0,-1,-1}
	vizinhosTarryR:= []int{0,0,-1,0,0}
	vizinhosTarryS:= []int{-1,-1,0,-1,0}
	vizinhosTarryT:= []int{-1,-1,0,0,-1}
	var primeiroTokenRecebidoP bool
	primeiroTokenRecebidoP=false
	var primeiroTokenRecebidoQ bool
	primeiroTokenRecebidoQ=false
	var primeiroTokenRecebidoR bool
	primeiroTokenRecebidoR=false
	var primeiroTokenRecebidoS bool
	primeiroTokenRecebidoS=false
	var primeiroTokenRecebidoT bool
	primeiroTokenRecebidoT=false
	go callbackTarry(chans, &vizinhosTarryP, &vizinhosTarryQ, &vizinhosTarryR, &vizinhosTarryS, &vizinhosTarryT, &primeiroTokenRecebidoP, &primeiroTokenRecebidoQ, &primeiroTokenRecebidoR, &primeiroTokenRecebidoS, &primeiroTokenRecebidoT)
	go processP(chans, &vizinhosTarryP, &primeiroTokenRecebidoP)
	go processQ(chans, &vizinhosTarryQ, &primeiroTokenRecebidoQ)
	go processR(chans, &vizinhosTarryR, &primeiroTokenRecebidoR)
	go processS(chans, &vizinhosTarryS, &primeiroTokenRecebidoS)
	go processT(chans, &vizinhosTarryT, &primeiroTokenRecebidoT)

	time.Sleep(5 * time.Second)
	fmt.Printf("\n")
}
