import { HeartTwoTone, SmileTwoTone } from '@ant-design/icons';
import { PageContainer } from '@ant-design/pro-components';
import { Alert, Card, Typography } from 'antd';
import React from 'react';

const Index: React.FC = () => {
  return (
    <PageContainer>
      <Card>
        <Alert
          type="success"
          showIcon
          banner
          style={{
            margin: -12,
            marginBottom: 48,
          }}
        />
        <Typography.Title level={2} style={{ textAlign: 'center' }}>
          <SmileTwoTone /> Ant Design Pro{' '}
          <HeartTwoTone twoToneColor="#eb2f96" /> You
        </Typography.Title>
      </Card>

    </PageContainer>
  );
};

export default Index;
