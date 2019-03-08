import React from 'react'

import { BrowserRouter as Router, Route } from 'react-router-dom'

// THEME
import MuiThemeProvider from '@material-ui/core/styles/MuiThemeProvider'
import muiTheme from './theme/muiTheme'

import App from './App'

const router = () => (
  <Router>
    <div>
      <Route
        path="/"
        exact
        component={() => (
          <MuiThemeProvider theme={muiTheme}>
            <App />
          </MuiThemeProvider>
        )}
      />
      <Route path="/about" component={() => <div>Hello World</div>} />
    </div>
  </Router>
)

export default router
