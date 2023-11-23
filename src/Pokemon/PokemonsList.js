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

const fetchPokemonDetail = async (url) => {
  try {
    const response = await fetch(url);
    const pokemonData = response.data;

    const htmlContent = `
        <div class="content-details">
          <a class="content-back" href="./index.html">&#8249;</a>
          <h1 class="content-details-heading">${pokemonData.name}</h1>
          <img class="content-details-img" src="${
            pokemonData.sprites.front_default
          }" alt="${pokemonData.name}"/>

          <table>
            <tr>
                <th>Weight</th>
                <th>Height</th>
                <th>Types</th>
            </tr>
            <tr>
                <td>${pokemonData.weight}</td>
                <td>${pokemonData.height}</td>
                <td>${pokemonData.types
                  .map((types) => types.type.name)
                  .join(', ')}</td>
            </tr>
          </table>
        <div>
        `;

    const detailElement = document.createElement('div');
    detailElement.innerHTML = htmlContent;

    const resultElement = document.getElementByClass('pokedex');
    resultElement.innerHTML = '';

    resultElement.appendChild(detailElement);
  } catch (error) {
    console.log(`Terjadi Error Pada ${error}`);
  }
};