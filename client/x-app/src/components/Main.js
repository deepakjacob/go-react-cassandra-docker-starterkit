import React from "react";
import PropTypes from "prop-types";

// @material-ui/core
import withStyles from "@material-ui/core/styles/withStyles";

import CssBaseline from "@material-ui/core/CssBaseline";
import Dashboard from "./Content/Dashboard"

import dashboardStyle from "../assets/jss/material-dashboard-react/layouts/dashboardStyle";

const Main = props => {
  const { classes, rest } = props;
  return (
    <div>
      <CssBaseline />
      <div className={classes.wrapper}>
        <div className={classes.mainPanel}>
          <div className={classes.content}>
            <Dashboard {...rest}/>
          </div>
        </div>
      </div>
    </div>
  );
}

Dashboard.propTypes = {
  classes: PropTypes.object.isRequired
};

export default withStyles(dashboardStyle)(Main);
