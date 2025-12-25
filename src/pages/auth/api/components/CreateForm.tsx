import { PlusOutlined } from '@ant-design/icons';
import {
  type ActionType,
  ModalForm, ProForm,ProFormSwitch,
  ProFormText,
  ProFormTextArea,
} from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Button, message } from 'antd';
import React, { FC } from 'react';
import { createApi }from "@/pages/auth/api/service/service";
import { Api } from  "@/pages/auth/api/service/data";
interface CreateFormProps {
  reload?: ActionType['reload'];
}

const CreateForm: FC<CreateFormProps> = (props) => {
  const { reload } = props;

  const [messageApi, contextHolder] = message.useMessage();


  const { run, loading } = useRequest(createApi, {
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
        modalProps={{ okButtonProps: { loading } }}
        onFinish={async (value) => {
          value.status = value.status?1:0;
          await run({ data: value as Api });

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
            name="title"
            label="标题"
            placeholder="请输入标题"
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
            width="md"
            name="icon"
            label="图标"
            placeholder="请输入图标"
            rules={[
              {
                required: true,
                message: '不能为空',
              },
            ]}
          />
          <ProFormText
            width="md"
            name="path"
            label="路由路径"
            placeholder="请输入路由路径"
            rules={[
              {
                required: true,
                message: '不能为空',
              },
            ]}
          />

          <ProFormText
            width="md"
            name="component"
            label="组件路径"
            placeholder="请输入组件路径"
            rules={[
              {
                required: true,
                message: '不能为空',
              },
            ]}
          />

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
