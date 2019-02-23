import React from 'react';

// THEME
import MuiThemeProvider from '@material-ui/core/styles/MuiThemeProvider';
import muiTheme from './theme/muiTheme';

// COMPONENTS
import Main from './components/Main'

const App = (props) => (
  <MuiThemeProvider theme={muiTheme}>
    <Main/>
  </MuiThemeProvider>
)

export default App;
