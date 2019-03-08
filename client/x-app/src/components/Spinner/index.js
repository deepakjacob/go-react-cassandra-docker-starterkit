import React from 'react'
import PropTypes from 'prop-types'
import { withStyles } from '@material-ui/core/styles'
import LinearProgress from '@material-ui/core/LinearProgress'

const styles = {
  root: {
    flexGrow: 1
  }
}

const LinearIndeterminate = props => {
  const { classes, options: { container, component } } = props
  return (
    <div className={classes.root}>
      {container && <LinearProgress color="primary" />}
      {component && <LinearProgress color="secondary" />}
    </div>
  )
}

LinearIndeterminate.propTypes = {
  classes: PropTypes.object.isRequired
}

export default withStyles(styles)(LinearIndeterminate)
