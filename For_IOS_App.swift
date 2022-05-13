
func PlaceBombe(M: inout [[Int]]) -> [[Int]]{
    // var Emplacement = 0
        for i in 0...M.count-1{
          
            for j in 0...M[0].count-1{
               
                if M[i][j] == 9{
                    
                    if i == 0 && j == 0{//[0][0] Exemple : [X . . . .]

                        if M[0][1] != 9{
                            M[0][1] += 1
                        }
                        if M[1][0] != 9{
                            M[1][0] += 1
                        }
                        if M[1][1] != 9{
                            M[1][1] += 1
                        }    

                    }else if i == 0 && j == M[0].count-1{//[0][max] Exemple : [. . . . X]

                        if M[0][M[0].count-2] != 9 {
                            M[0][M[0].count-2] += 1
                        }
                        if M[1][M[0].count-1] != 9 {
                            M[1][M[0].count-1] += 1
                        }
                        if M[1][M[0].count-2] != 9 {
                            M[1][M[0].count-2] += 1
                        }

                    }else if i == M.count-1 && j == 0{ //[max][0] Exemple : [X . . . .]

                        if M[M.count-1][1] != 9{
                            M[M.count-1][1] += 1
                        }
                        if M[M.count-2][0] != 9{
                            M[M.count-2][0] += 1
                        }
                        if M[M.count-2][1] != 9{
                            M[M.count-2][1] += 1
                        }

                    }else if i == M.count-1 && j == M[0].count-1{//[max][max] Exemple : [. . . . X]

                        if M[M.count-1][M[0].count-2] != 9{
                            M[M.count-1][M[0].count-2] += 1
                        }
                        if M[M.count-2][M[0].count-2] != 9{
                            M[M.count-2][M[0].count-2] += 1
                        }
                        if M[M.count-2][M[0].count-1] != 9{
                            M[M.count-2][M[0].count-1] += 1
                        }

                    }else if i == 0{

                        if j != 0 && j != M[0].count-1{ //In [0][...] Exemple : [. X X X .]

                            if M[0][j-1] != 9{
                                M[0][j-1] += 1
                            }
                            if M[0][j+1] != 9{
                                M[0][j+1] += 1
                            }
                            if M[1][j-1] != 9{
                                M[1][j-1] += 1
                            }
                            if M[1][j] != 9{
                                M[1][j] += 1
                            }
                            if M[1][j+1] != 9{
                                M[1][j+1] += 1
                            }

                        }
                    }else if i == M.count-1{//In [MAX][...] Exemple : [. X X X .]
                        
                        if j != 0 && j != M[0].count-1{

                            if M[M.count-1][j-1] != 9{
                                M[M.count-1][j-1] += 1
                            }
                            if M[M.count-1][j+1] != 9{
                                M[M.count-1][j+1] += 1
                            }
                            if M[M.count-2][j-1] != 9{
                                M[M.count-2][j-1] += 1
                            }
                            if M[M.count-2][j] != 9{
                                M[M.count-2][j] += 1
                            }
                            if M[M.count-2][j+1] != 9{
                                M[M.count-2][j+1] += 1
                            }
                            
                        }

                    }else if i != 0 && i != M.count-1{//Border

                        if j == 0{// Border left

                            if M[i-1][0] != 9{
                                M[i-1][0] += 1
                            }
                            if M[i-1][1] != 9{
                                M[i-1][1] += 1
                            }
                            if M[i][1] != 9{
                                M[i][1] += 1
                            }
                            if M[i+1][0] != 9{
                                M[i+1][0] += 1
                            }
                            if M[i+1][1] != 9{
                                M[i+1][1] += 1
                            }
                            
                        }else if j == M[0].count-1{// Border right

                            if M[i-1][M[0].count-1] != 9{
                                M[i-1][M[0].count-1] += 1
                            }
                            if M[i-1][M[0].count-2] != 9{
                                M[i-1][M[0].count-2] += 1
                            }
                            if M[i][M[0].count-2] != 9{
                                M[i][M[0].count-2] += 1
                            }
                            if M[i+1][M[0].count-2] != 9{
                                M[i+1][M[0].count-2] += 1
                            }
                            if M[i+1][M[0].count-1] != 9{
                                M[i+1][M[0].count-1] += 1
                            }

                        }else{//The rest

                            if M[i-1][j-1] != 9{
                                M[i-1][j-1] += 1
                            }
                            if M[i-1][j] != 9{
                                M[i-1][j] += 1
                            }
                            if M[i-1][j+1] != 9{
                                M[i-1][j+1] += 1
                            }
                            if M[i][j-1] != 9{
                                M[i][j-1] += 1
                            }
                            if M[i][j+1] != 9{
                                M[i][j+1] += 1
                            }
                            if M[i+1][j-1] != 9{
                                M[i+1][j-1] += 1
                            }
                            if M[i+1][j] != 9{
                                M[i+1][j] += 1
                            }
                            if M[i+1][j+1] != 9{
                                M[i+1][j+1] += 1
                            }
                        }
                    }
                }
            }
        }
    return M
}

//Define Bombe
func Define_Map(Difficulty: Int,M: inout [[Int]]) -> ([[Int]],Int) {
    
    var count = 0

    for i in 0...M.count-1{
        for j in 0...M[i].count-1{
            var Rdn_Bombe = Int.random(in: 0..<Difficulty)
            if Rdn_Bombe == 1{
               M[i][j] = 9
               count += 1
            }
        }
    }
    return (M, count)
}

//Create the Map with Size
func Create_Map(x: Int,y: Int ) -> [[Int]]{
    
    var Map = [[Int]]()
    
    for i in 0..<x{
        var vide = [Int]()
        for _ in 0..<y{
            vide.append(0)
        }
        Map.insert(vide,at:i)
    }
    return Map
}

//Func Main
func main(){
    let Size_x = 40 // Change to var
    let Size_y = 40 // Change to var
    var Difficulty = 5

    
    var Map = Create_Map(x:Size_x,y:Size_y)
    var (Map_end, Bombes) = Define_Map(Difficulty:Difficulty, M: &Map)
    var Map_end_with_Bombe = PlaceBombe(M: &Map_end)
    
    print("")
    for i in 0...Map_end_with_Bombe.count-1{
       print(Map_end_with_Bombe[i])
    }
    print("\nbombes: ",Bombes)

}

main()
