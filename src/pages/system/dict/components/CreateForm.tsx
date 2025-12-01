import { PlusOutlined } from '@ant-design/icons';
import {
  type ActionType,
  ModalForm, ProForm,
  ProFormText,
} from '@ant-design/pro-components';
import {  useRequest } from '@umijs/max';
import { Button, message } from 'antd';
import React, { FC } from 'react';
import { Dictht,Dict } from  "@/pages/system/dict/service/data";
import {createDict}  from "@/pages/system/dict/service/service";

interface CreateFormProps {
  reload?: ActionType['reload'];
}

const CreateForm: FC<CreateFormProps> = (props) => {
  const { reload } = props;

  const [messageApi, contextHolder] = message.useMessage();


  const { run, loading } = useRequest(createDict, {
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
       title='新建'
        trigger={
          <Button type="primary" icon={<PlusOutlined />}>
            新建
          </Button>
        }
        width="400px"
        modalProps={{ okButtonProps: { loading } }}
        onFinish={async (value) => {
          await run({ data: value as Dict });

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
            name="desc"
            label="概略"
            placeholder="请输入"
          />

          <ProFormText
            width="md"
            name="status"
            label="状态"
          />

        </ProForm.Group>
      </ModalForm>
    </>
  );
};

export default CreateForm;
