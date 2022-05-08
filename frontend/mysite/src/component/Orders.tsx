import * as React from 'react';
import Link from '@mui/material/Link';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Title from './Title';
import { DataGrid, GridColDef, GridRowsProp } from '@mui/x-data-grid';

// Generate Order Data
function createData(
  id: number,
  date: string,
  name: string,
  shipTo: string,
  paymentMethod: string,
  amount: number,
) {
  return { id, date, name, shipTo, paymentMethod, amount };
}

const rows = [
  createData(
    0,
    '16 Mar, 2019',
    'Elvis Presley',
    'Tupelo, MS',
    'VISA ⠀•••• 3719',
    312.44,
  ),
  createData(
    1,
    '16 Mar, 2019',
    'Paul McCartney',
    'London, UK',
    'VISA ⠀•••• 2574',
    866.99,
  ),
  createData(2, '16 Mar, 2019', 'Tom Scholz', 'Boston, MA', 'MC ⠀•••• 1253', 100.81),
  createData(
    3,
    '16 Mar, 2019',
    'Michael Jackson',
    'Gary, IN',
    'AMEX ⠀•••• 2000',
    654.39,
  ),
  createData(
    4,
    '15 Mar, 2019',
    'Bruce Springsteen',
    'Long Branch, NJ',
    'VISA ⠀•••• 5919',
    212.79,
  ),
];

function preventDefault(event: React.MouseEvent) {
  event.preventDefault();
}

class Orders extends React.Component<any, any> {
  render(): React.ReactNode {
    return (
      <React.Fragment>
        <Title>Recent Orders</Title>
        <Table size="small">
          <TableHead>
            <TableRow>
              <TableCell>Date</TableCell>
              <TableCell>Name</TableCell>
              <TableCell>Ship To</TableCell>
              <TableCell>Payment Method</TableCell>
              <TableCell align="right">Sale Amount</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row) => (
              <TableRow key={row.id}>
                <TableCell>{row.date}</TableCell>
                <TableCell>{row.name}</TableCell>
                <TableCell>{row.shipTo}</TableCell>
                <TableCell>{row.paymentMethod}</TableCell>
                <TableCell align="right">{`$${row.amount}`}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        <Link color="primary" href="#" onClick={preventDefault} sx={{ mt: 3 }}>
          See more orders
        </Link>
      </React.Fragment>
    );
  }
}

class OrdersX extends React.Component<any, any> {

  constructor(props: any) {
    super(props)
  }
  render(): React.ReactNode {
    const rs: GridRowsProp = rows.map((r, index) => ({
      id: index, col1: r.date, col2: r.name, col3: r.shipTo, col4: r.paymentMethod, col5: r.amount
    }))
    console.log(rs)
    const cols: GridColDef[] = [
      { field: 'col1', headerName: 'column 1', width: 150 ,headerClassName: 'super-app-theme--header',
      headerAlign: 'center',},
      { field: 'col2', headerName: 'column 2', width: 150 },
      { field: 'col3', headerName: 'column 3', width: 150 },
      { field: 'col4', headerName: 'column 4', width: 150 },
      { field: 'col5', headerName: 'column 5', width: 150 },
    ]
    return (<>
      <DataGrid rows={rs} columns={cols}></DataGrid>
    </>)
  }
}

export default OrdersX 
