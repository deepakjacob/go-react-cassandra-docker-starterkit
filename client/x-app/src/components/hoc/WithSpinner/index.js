import React, { useState } from 'react'
import Spinner from '../../Spinner'

const withSpinner = options => {
  return WrappedComponent => props => {
    const [loading, setLoading] = useState(true)
    setTimeout(() => setLoading(false), 3000)
    return loading ? (
      <div>
        <Spinner options={options} />
        <WrappedComponent {...props} />
      </div>
    ) : (
      <WrappedComponent {...props} />
    )
  }
}

export default withSpinner
