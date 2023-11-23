import React from 'react';
import Pokemon from './Pokemon';

export const listPokemons = ({ pokemons, onClick, buttonLabel }) =>
  pokemons.map((pokemon) => (
    <Pokemon key={pokemon.name} pokemon={pokemon} onClick={onClick} buttonLabel={buttonLabel} />
  ));