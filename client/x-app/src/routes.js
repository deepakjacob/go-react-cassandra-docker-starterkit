import React from 'react'

import { BrowserRouter as Router, Route } from 'react-router-dom'
import App from './App'

const router = () => (
  <Router>
    <div>
      <Route path="/" exact component={App} />
      <Route path="/about" component={() => <div>Hello World</div>} />
    </div>
  </Router>
)

export default router
