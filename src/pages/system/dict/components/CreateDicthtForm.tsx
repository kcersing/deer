import { PlusOutlined } from '@ant-design/icons';
import {
  type ActionType,
  ModalForm, ProForm,
  ProFormText,
  ProFormTextArea,ProFormSwitch
} from '@ant-design/pro-components';
import {  useRequest } from '@umijs/max';
import { Button, message } from 'antd';
import React, { FC } from 'react';

import { Dictht,Dict } from  "@/pages/system/dict/service/data";
import {createDictht}  from "@/pages/system/dict/service/service";

interface CreateFormProps {
  reload?: ActionType['reload'];
}

const CreateDicthtForm: FC<CreateFormProps> = (props) => {
  const { reload,dictId } = props;

  const [messageApi, contextHolder] = message.useMessage();


  const { run, loading } = useRequest(createDictht, {
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
        width="800px"
        modalProps={{ okButtonProps: { loading } }}
        onFinish={async (value) => {
          value.dictId = dictId;
          value.status = value.status?1:0;
           await run({ data: value as Dictht });
          return true;
        }}
      >
        <ProForm.Group>

          <ProFormText
            width="md"
            name="title"
            label="标题"
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
          width="md"
          name="value"
          label="有效值"
          placeholder="请输入"
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
            checkedChildren="开启"
            unCheckedChildren="关闭"
            initialValue="开启"
          />

        </ProForm.Group>
      </ModalForm>
    </>
  );
};

export default CreateDicthtForm;
