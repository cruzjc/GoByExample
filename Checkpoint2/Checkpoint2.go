package main

import(
	"fmt"
	"time"
)


func main(){
	const player1Name string = "player1"
	const player2Name string = "player2"
	startingHealthPlayer1:=100
	startingHealthPlayer2:=100
	currentHealthPlayer1:=startingHealthPlayer1
	currentHealthPlayer2:=startingHealthPlayer2
	player1Attack:=10
	player2Attack:=10



	isPlayer1Alive:=currentHealthPlayer1>0
	isPlayer2Alive:=currentHealthPlayer2>0
	areAllPlayersAlive:=isPlayer1Alive||isPlayer2Alive

	
	numberOfPlayers:=2
	playersInGame:=make([]string,0)
	
	fmt.Println("Loading Players")
	playersInGame= append(playersInGame,player1Name,player2Name)
	fmt.Println(player1Name + " has joined")
	fmt.Println(player2Name + " has joined")


	randomTurn:=(int)(time.Now().Unix())%(numberOfPlayers)+1
	for areAllPlayersAlive{

		
		switch randomTurn{
			case 1:
				fmt.Println("Player 1 attacks player 2")
				currentHealthPlayer2-=player1Attack
			case 2:
				fmt.Println("Player 2 attacks player 1")
				currentHealthPlayer1-=player2Attack
			default:
				fmt.Println("Invalid roll")
		}

		fmt.Println("Player 1 health: ", currentHealthPlayer1)
		fmt.Println("Player 2 health: ", currentHealthPlayer2)

		if randomTurn==numberOfPlayers{
			randomTurn=0
		}
		randomTurn+=1;


		isPlayer1Alive=currentHealthPlayer1>0
		isPlayer2Alive=currentHealthPlayer2>0
		areAllPlayersAlive=isPlayer1Alive&&isPlayer2Alive
	}

}