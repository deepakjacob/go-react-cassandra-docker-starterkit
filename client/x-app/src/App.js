import React from 'react'

// util
import { compose } from 'recompose'

// HOCs
import withSpinner from './components/hoc/WithSpinner'

// COMPONENTS
import Main from './components/Main'

var withLoader = WrappedComponent => props => {
  const mergedProps = {
    ...props,
    ...{ loadingText: 'Loading....', loading: true }
  }
  return <WrappedComponent {...mergedProps} />
}
var app = props => <Main {...props} />
export default compose(withLoader, withSpinner({ container: true }))(app)
