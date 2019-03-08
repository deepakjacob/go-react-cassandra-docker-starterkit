import React, { useState } from 'react'

// THEME
import MuiThemeProvider from '@material-ui/core/styles/MuiThemeProvider'
import muiTheme from './theme/muiTheme'

import { compose, branch } from 'recompose'

// COMPONENTS
import Main from './components/Main'
import Spinner from './components/Spinner'

const App = props => (
  <MuiThemeProvider theme={muiTheme}>
    <Main />
  </MuiThemeProvider>
)

const withSpinner = WrappedComponent => props => {
  const [loading, setLoading] = useState(true)
  setTimeout(() => setLoading(false), 3000)
  return loading ? <Spinner loadingText="Loading..." /> : <WrappedComponent />
}
const withPropLoader = WrappedComponent => props => {
  const mergedProps = {
    ...props,
    ...{ loadingText: 'Loading....' }
  }
  return <WrappedComponent {...mergedProps} />
}

export default compose(withPropLoader, withSpinner)(App)
