import React, { useContext, useEffect } from 'react';
import { PokemonContext } from './PokemonContext';
import { listPokemons } from './listPokemons';

const url = 'https://pokeapi.co/api/v2/pokemon';

const PokemonsList = () => {
  const { pokemons, capture, addPokemons } = useContext(PokemonContext);

  useEffect(() => {
    const fetchPokemons = async () => {
      // const response = await fetch(url);
      // const data = await response.json();
      // addPokemons(data.results);

      const promises = [];
      for(let i = 1; i <= 20; i++) {
        promises.push(fetch(url+`/${i}`)
        .then(res => res.json()));
      }

      Promise.all(promises).then(result => {
        const pokemon = result.map(data => ({
          id: data.id,
          name: data.name,
          image: data.sprites['front_default'],
        }));
        addPokemons(pokemon);
      })
    };

    fetchPokemons();
    // eslint-disable-next-line
  }, []);

  return (
    <div className="pokemons-list">
      <h2>Pokemon List</h2>
        <ul>
          {listPokemons({
            pokemons,
            onClick: capture,
            buttonLabel: 'Catch',
          })}
        </ul>
    </div>
  );
};

export default PokemonsList;
