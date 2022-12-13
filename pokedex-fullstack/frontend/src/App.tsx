import { useContext, useEffect } from 'react'
// CONTEXT
import PokemonContext from './context/PokemonContext'

function App() {
  const { getPokemons } = useContext(PokemonContext)

  useEffect(() => {
    getPokemons()
  }, [])

  return <div className=''>Pokedex</div>
}

export default App
