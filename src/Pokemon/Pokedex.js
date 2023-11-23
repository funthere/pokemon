import React, { useContext } from 'react';
import { PokemonContext } from './PokemonContext';
import { listPokemons } from './listPokemons';

const Pokedex = () => {
  const { capturedPokemons, release } = useContext(PokemonContext);

  return (
    <div className="pokedex">
      <h2>My Pokemon List</h2>
        <ul>
          {listPokemons({
            pokemons: capturedPokemons,
            onClick: release,
            buttonLabel: 'Release',
          })}
        </ul>
    </div>
  );
};

export default Pokedex;