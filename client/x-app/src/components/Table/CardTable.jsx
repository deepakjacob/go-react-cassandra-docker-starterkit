import React from 'react'

import { compose } from 'recompose'
// core components
import Card from '../Card/Card'
import Table from '../Table/Table.jsx'
import CardBody from '../Card/CardBody'
import CardHeader from '../Card/CardHeader'

const _CardTable = props => (
  <Table
    tableHeaderColor="warning"
    tableHead={['ID', 'Name', 'Salary', 'Country']}
    tableData={[
      ['1', 'Dakota Rice', '$36,738', 'Niger'],
      ['2', 'Minerva Hooper', '$23,789', 'Curaçao'],
      ['3', 'Sage Rodriguez', '$56,142', 'Netherlands'],
      ['4', 'Philip Chaney', '$38,735', 'Korea, South']
    ]}
  />
)

const withHeader = WrappedComponent => props => (
  <Card>
    <CardHeader color="warning">
      <h4 className={props.classes.cardTitleWhite}>Employees Stats</h4>
      <p className={props.classes.cardCategoryWhite}>
        New employees on 15th September, 2016
      </p>
    </CardHeader>
    <CardBody>
      <WrappedComponent {...props} />
    </CardBody>
  </Card>
)

const CardTable = compose(withHeader)(_CardTable)

export default CardTable
