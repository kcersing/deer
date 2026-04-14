import { List, Switch } from 'antd';
import React, { cloneElement, useCallback, useState } from 'react';
import { message } from 'antd';
import { BetaSchemaForm } from '@ant-design/pro-components';
import {changePassword} from "../service";
import {  useRequest } from '@umijs/max';
  type Password = {
    oldPassword: string;
    password: string;

  }
  const columns: ProFormColumnsType<Password>[] = [
      {
        title: '旧密码',
        description: '旧密码',
        dataIndex: 'oldPassword',
        valueType: 'password',
        tooltip: '如忘记密码请联系管理员重置',
        formItemProps: {
          rules: [
            {
              required: true,
              message: '此项为必填项',
            },
          ],
        },
      },
      {
        title: '新密码',
        description: '新密码',
        dataIndex: 'newPassword',
        valueType: 'password',
        formItemProps: {
          rules: [
            {
              required: true,
              message: '此项为必填项',
            },
          ],
        },
      },
      {
        title: '确认密码',
        dataIndex: 'newPassword1',
        description: '确认密码',
        valueType: 'password',
        formItemProps: {
          rules: [
            {
              required: true,
              message: '此项为必填项',
            },
          ],
        },
      },
    ];



const ChangePassword: React.FC = () => {

  return (
    <div style={{padding: 24}}>
      <BetaSchemaForm<DataItem>
        layoutType="Form"

        onFinish={async (values) => {

          if ( values.newPassword === values.oldPassword) {
            message.error('新旧不能相同！');
          }
          if ( values.newPassword !== values.newPassword1) {
            message.error('新密码不一致！');
          }
          values.oldPassword  = values.oldPassword;
          values.password  = values.newPassword;

          const msg = await changePassword({...values});
          if (msg.code == 0) {
            message.success(msg.message);
          }else {
            message.error(msg.message);
          }
          return

        }}

        columns={columns}
      />
    </div>
  );
};


export default ChangePassword;
