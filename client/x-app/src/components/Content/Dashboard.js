import React from "react";
import PropTypes from "prop-types";

// @material-ui/core
import Icon from "@material-ui/core/Icon";
import withStyles from "@material-ui/core/styles/withStyles";
// @material-ui/icons
import Store from "@material-ui/icons/Store";
import Update from "@material-ui/icons/Update";
import Warning from "@material-ui/icons/Warning";
import DateRange from "@material-ui/icons/DateRange";
import LocalOffer from "@material-ui/icons/LocalOffer";
import AccessTime from "@material-ui/icons/AccessTime";
import ArrowUpward from "@material-ui/icons/ArrowUpward";
import Accessibility from "@material-ui/icons/Accessibility";

// core components
import Card from "../Card/Card";
import CardBody from "../Card/CardBody";
import CardIcon from "../Card/CardIcon";
import GridItem from "../Grid/GridItem";
import Danger from "../Typography/Danger";
import CardHeader from "../Card/CardHeader";
import CardFooter from "../Card/CardFooter";
import GridContainer from "../Grid/GridContainer";
import CardTable from '../Table/CardTable.jsx'
import dashboardStyle from "../../assets/jss/material-dashboard-react/views/dashboardStyle";

const DailyStatsCard = (props) => <CardTable {...props}/>

const SmallCards = props => {
  const { classes } = props;
  return (
    <GridContainer>
      <GridItem xs={12} sm={6} md={3}>
        <Card>
          <CardHeader color="warning" stats icon>
            <CardIcon color="warning">
              <Icon>content_copy</Icon>
            </CardIcon>
            <p className={classes.cardCategory}>Used Space</p>
            <h3 className={classes.cardTitle}>
              49/50 <small>GB</small>
            </h3>
          </CardHeader>
          <CardFooter stats>
            <div className={classes.stats}>
              <Danger>
                <Warning />
              </Danger>
              <a href="#pablo" onClick={e => e.preventDefault()}>
                Get more space
              </a>
            </div>
          </CardFooter>
        </Card>
      </GridItem>
      <GridItem xs={12} sm={6} md={3}>
        <Card>
          <CardHeader color="success" stats icon>
            <CardIcon color="success">
              <Store />
            </CardIcon>
            <p className={classes.cardCategory}>Revenue</p>
            <h3 className={classes.cardTitle}>$34,245</h3>
          </CardHeader>
          <CardFooter stats>
            <div className={classes.stats}>
              <DateRange />
              Last 24 Hours
            </div>
          </CardFooter>
        </Card>
      </GridItem>
      <GridItem xs={12} sm={6} md={3}>
        <Card>
          <CardHeader color="danger" stats icon>
            <CardIcon color="danger">
              <Icon>info_outline</Icon>
            </CardIcon>
            <p className={classes.cardCategory}>Fixed Issues</p>
            <h3 className={classes.cardTitle}>75</h3>
          </CardHeader>
          <CardFooter stats>
            <div className={classes.stats}>
              <LocalOffer />
              Tracked from Github
            </div>
          </CardFooter>
        </Card>
      </GridItem>
      <GridItem xs={12} sm={6} md={3}>
        <Card>
          <CardHeader color="info" stats icon>
            <CardIcon color="info">
              <Accessibility />
            </CardIcon>
            <p className={classes.cardCategory}>Followers</p>
            <h3 className={classes.cardTitle}>+245</h3>
          </CardHeader>
          <CardFooter stats>
            <div className={classes.stats}>
              <Update />
              Just Updated
            </div>
          </CardFooter>
        </Card>
      </GridItem>
    </GridContainer>
  );
};

const GraphCard = props => {
  const { classes } = props;
  return (
    <Card chart>
      <CardHeader color="success" />
      <CardBody>
        <h4 className={classes.cardTitle}>Daily Sales</h4>
        <p className={classes.cardCategory}>
          <span className={classes.successText}>
            <ArrowUpward className={classes.upArrowCardCategory} /> 55%
          </span>{" "}
          increase in today sales.
        </p>
      </CardBody>
      <CardFooter chart>
        <div className={classes.stats}>
          <AccessTime /> updated 4 minutes ago
        </div>
      </CardFooter>
    </Card>
  );
};
const Dashboard = props => {
  return (
    <div>
      <SmallCards {...props} />
      <GridContainer alignItems="center">
        <GridItem xs={12} sm={12} md={6}>
          <GraphCard {...props} />
        </GridItem>
        <GridItem xs={12} sm={12} md={6}>
          <GraphCard {...props} />
        </GridItem>
        <GridItem xs={12} sm={12} md={6}>
          <DailyStatsCard {...props} />
        </GridItem>
        <GridItem xs={12} sm={12} md={6}>
          <DailyStatsCard {...props} />
        </GridItem>
      </GridContainer>
    </div>
  );
};

Dashboard.propTypes = {
  classes: PropTypes.object.isRequired
};

export default withStyles(dashboardStyle)(Dashboard);
