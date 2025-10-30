import { history } from '@umijs/max';
import { Button, Card, Result } from 'antd';
import React from 'react';

const NoFoundPage: React.FC = () => (
  <Card variant="borderless">
    <Result
      status="404"
      title="404"
      subTitle="404"
      extra={
        <Button type="primary" onClick={() => history.push('/')}>
         跳转
        </Button>
      }
    />
  </Card>
);

export default NoFoundPage;
