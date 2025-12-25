import { PlusOutlined } from '@ant-design/icons';
import {
  type ActionType,
  ModalForm,
  ProForm,
  ProFormText,
  ProFormTextArea,
  ProFormTreeSelect,
  ProFormSwitch,
} from '@ant-design/pro-components';
import {  useRequest } from '@umijs/max';
import { Button, message ,TreeSelect} from 'antd';
import React, { FC } from 'react';

import { createRole }from "@/pages/auth/role/service/service";
import { Role } from  "@/pages/auth/role/service/data";

import { getMenuList } from "@/pages/auth/menu/service/service";
import { getApiList } from "@/pages/auth/api/service/service";

interface CreateFormProps {
  reload?: ActionType['reload'];
}

const CreateForm: FC<CreateFormProps> = (props) => {
  const { reload } = props;

  const [messageApi, contextHolder] = message.useMessage();

  const { run, loading } = useRequest(createRole, {
    manual: true,
    onSuccess: () => {
      messageApi.success('提交成功');
      reload?.();
    },
    onError: () => {
      messageApi.error('提交失败，请重试！');
    },
  });

  return (
    <>
      {contextHolder}
      <ModalForm
       title='新建菜单'
        trigger={
          <Button type="primary" icon={<PlusOutlined />}>
            新建
          </Button>
        }

        width="400px"
        modalProps={{
          okButtonProps: { loading },
          destroyOnHidden: true,
          onCancel: () => console.log('run'),
      }}
        onFinish={async (value) => {
          value.status = value.status?1:0;
          await run({ data: value as Role });

          return true;
        }}
      >

        <ProForm.Group>
          <ProFormText
            width="md"
            name="name"
            label="名称"
            tooltip="最长为 24 位"
            placeholder="请输入名称"
            rules={[
              {
                required: true,
                message: '不能为空',
              },
            ]}
          />

          <ProFormText
            width="md"
            name="code"
            label="标识"
            placeholder="请输入"
            rules={[
              {
                required: true,
                message: '不能为空',
              },
            ]}
          />
          <ProFormText
            width="md"
            name="desc"
            label="简介"
            placeholder="请输入"
            rules={[
              {
                required: true,
                message: '不能为空',
              },
            ]}
          />

        </ProForm.Group>

        <ProForm.Group>
          <ProFormTreeSelect

            label="菜单权限"
            width={260}
            params={{current: 999, pageSize: 1}}
            request={(params)=>{
              return getMenuList({params}).then((res) => {return res.data})
            }}
            name="menus"
            fieldProps={{
              fieldNames: {
                label: 'name',
                value: 'id',
                children: 'children',
              },
              allowClear:true,
              treeCheckable: true,
              multiple: true,
              showCheckedStrategy: TreeSelect.SHOW_ALL,
              placeholder: '请选择',
            }}
          />
        </ProForm.Group>

        <ProForm.Group>
          <ProFormTreeSelect

            label="API权限"
            width={260}
            params={{current: 999, pageSize: 1}}
            request={(params)=>{
              return getApiList({params}).then((res) => {return res.data})
            }}
            name="apis"
            fieldProps={{
              fieldNames: {
                label: 'name',
                value: 'id',
                children: 'children',
              },
              allowClear:true,
              treeCheckable: true,
              multiple: true,
              showCheckedStrategy: TreeSelect.SHOW_ALL,
              placeholder: '请选择',
            }}
          />
        </ProForm.Group>

        <ProForm.Group>
          <ProFormSwitch
            name="status"
            width="md"
            label="状态"
            checkedChildren="有效"
            unCheckedChildren="无效"
          />

        </ProForm.Group>

      </ModalForm>
    </>
  );
};

export default CreateForm;
