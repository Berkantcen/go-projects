import React, { useState, createContext } from 'react'
import axios from 'axios'

const PokemonContext = createContext(
  {} as {
    pokemons: never[]
    setPokemons: React.Dispatch<React.SetStateAction<never[]>>
    getPokemons: () => Promise<void>
  }
)

type PokemonProps = {
  children: React.ReactNode
}

export const PokemonProvider = ({ children }: PokemonProps) => {
  const [pokemons, setPokemons] = useState([])

  const getPokemons = async () => {
    await axios
      .get('http://localhost:8080/pokemons')
      .then((res) => {
        setPokemons(res.data)
      })
      .catch((err) => {
        console.log(err)
      })
  }

  return (
    <PokemonContext.Provider value={{ pokemons, setPokemons, getPokemons }}>
      {children}
    </PokemonContext.Provider>
  )
}

export default PokemonContext
