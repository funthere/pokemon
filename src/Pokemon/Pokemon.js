import React from 'react';

export const Pokemon = ({ pokemon, onClick, buttonLabel }) => (
  <li key={pokemon.name}>
        <img alt={pokemon.name} src={pokemon.image}></img>
        {pokemon.name}
        <br></br>
        <button onClick={onClick(pokemon)}>{buttonLabel}</button>
  </li>
);

export default Pokemon;