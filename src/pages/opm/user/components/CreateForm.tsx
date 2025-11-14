import { PlusOutlined } from '@ant-design/icons';
import {
  type ActionType,
  ModalForm, ProForm,
  ProFormText,
  ProFormTextArea,
  ProFormSwitch,
  ProFormDatePicker,
  ProFormUploadButton,ProFormRadio
} from '@ant-design/pro-components';
import {  useRequest } from '@umijs/max';
import { Button, message } from 'antd';
import React, { FC,useState } from 'react';

import {createUser} from "@/services/ant-design-pro/user";
// 导入 Slate 编辑器工厂。
import { createEditor } from 'slate'
// 导入 Slate 组件和 React 插件。
import { Slate, Editable, withReact } from 'slate-react'

interface CreateFormProps {
  reload?: ActionType['reload'];
}

const CreateForm: FC<CreateFormProps> = (props) => {
  const { reload } = props;

  const [messageApi, contextHolder] = message.useMessage();

  const [value, setValue] = useState('');
  const { run, loading } = useRequest(createUser, {
    manual: true,
    onSuccess: () => {
      messageApi.success('提交成功');
      reload?.();
    },
    onError: () => {
      messageApi.error('提交失败，请重试！');
    },
  });

  const [editor] = useState(() => withReact(createEditor()))
// 添加初始化值。



  const initialValue =  [{
    type: 'paragraph',
    children: [{ text: 'Default content' }]
  }];


  return (
    <>
      {contextHolder}
      <ModalForm
       title='新建'
        trigger={
          <Button type="primary" icon={<PlusOutlined />}>
            新建
          </Button>
        }
        width="800px"
        modalProps={{ okButtonProps: { loading } }}
        onFinish={async (value) => {
          await run({ data: value as API.User });
          return true;
        }}
      >


        {/*<ProCard style={{ marginBlockStart: 8 }} gutter={8} title="指定宽度px">*/}
        {/*  <ProCard*/}
        {/*    colSpan={{*/}
        {/*      xs: '50px',*/}
        {/*      sm: '100px',*/}
        {/*      md: '200px',*/}
        {/*      lg: '300px',*/}
        {/*      xl: '400px',*/}
        {/*    }}*/}
        {/*    layout="center"*/}
        {/*    bordered*/}
        {/*  >*/}
        {/*    Col*/}
        {/*  </ProCard>*/}
        {/*  <ProCard layout="center" bordered>*/}
        {/*    Auto*/}
        {/*  </ProCard>*/}
        {/*</ProCard>*/}



        <ProForm.Group>
          <ProFormUploadButton
            max={1}
            width="sm"
            name="avatar"
            label="头像"
            fieldProps={{
              name: 'file',
              listType: 'picture-card',
            }}
            action="/upload.do"
            extra="头像备注"
          />
        </ProForm.Group>
        {/*number | "xs" | "sm" | "md" | "lg" | "xl"*/}
        <ProForm.Group>
          <ProFormText
            width="sm"
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
            width="sm"
            name="username"
            label="登录账号"
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
          <ProFormText
            width="sm"
            name="mobile"
            label="手机号"
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

          <ProFormRadio.Group
            width="sm"
            name="gender"
            label="性别"
            options={[
              { value: 1, label: '男' },
              { value: 2, label: '女' },
              { value: 3, label: '保密 ' },
            ]}
          />

          <ProFormDatePicker
            width="sm"
            name="birthday"
            label="出生日期"
          />
          <ProFormSwitch
            width="sm"
            name="status"
            label="状态"
            checkedChildren="开启"
            unCheckedChildren="关闭"
          />

        </ProForm.Group>
        <ProForm.Group>
          <Slate editor={editor} initialValue={initialValue }>
            <Editable />
          </Slate>
        </ProForm.Group>
      </ModalForm>
    </>
  );
};
export default CreateForm;
