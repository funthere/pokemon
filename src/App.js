import React from 'react';

import { PokemonProvider } from './Pokemon/PokemonContext';
import PokemonsList from './Pokemon/PokemonsList';
import Pokedex from './Pokemon/Pokedex';

const App = () => (
    <PokemonProvider>
      <div className="main">
        <PokemonsList />
        <Pokedex />
      </div>
    </PokemonProvider>
);

export default App;