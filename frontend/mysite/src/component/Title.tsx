import * as React from 'react';
import Typography from '@mui/material/Typography';

interface TitleProps {
  children?: React.ReactNode;
}

class Title extends React.Component<TitleProps,any> {
  render(): React.ReactNode {
    return (
      <Typography component="h2" variant="h6" color="primary" gutterBottom>
        {this.props.children}
      </Typography>
    );    
  }
}

export default Title;
