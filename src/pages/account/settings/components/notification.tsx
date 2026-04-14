import { List, Switch } from 'antd';
import React from 'react';

import { message } from 'antd';
import { BetaSchemaForm } from '@ant-design/pro-components';



type DataItem = {
  name: string;
  state: string;
  title: string;
}
const columns: ProFormColumnsType<DataItem>[] = [
  {
    title: '用户消息',
    description: '其他用户的消息将以站内信的形式通知',
    dataIndex: 'userMessage',
    valueType: 'switch',
    formItemProps: {
      extra: '其他用户的消息将以站内信的形式通知',
    },
    fieldProps: {
      initialValue: '开启',
      checkedChildren: '开启',
      unCheckedChildren: '关闭',
    },
  },
  {
    title: '系统消息',
    description: '系统消息将以站内信的形式通知',
    dataIndex: 'systemMessage',
    valueType: 'switch',
    formItemProps: {
      extra: '系统消息将以站内信的形式通知',
    },
    fieldProps: {
      initialValue: '开启',
      checkedChildren: '开启',
      unCheckedChildren: '关闭',
    },
    width: 'md',
    colProps: {
      xs: 12,
      md: 20,
    },
  },

];


const NotificationView: React.FC = () => {
  return (
    <div style={{padding: 24}}>
      <BetaSchemaForm<DataItem>
        layoutType="Form"
        onFinish={async (values) => {

          console.log(values);

        }}
        columns={columns}
      />
    </div>
  );
};


export default NotificationView;
