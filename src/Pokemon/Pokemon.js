import React from 'react';

export const Pokemon = ({ pokemon, onClick, buttonLabel }) => (
  <li key={pokemon.name}>
        <img alt={pokemon.name} src={pokemon.image}></img>
        <span className="pokemon-name">{pokemon.name}</span>
        <br></br>
        <button onClick={onClick(pokemon)}>{buttonLabel}</button>
  </li>
);

export default Pokemon;