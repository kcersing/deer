import { PlusOutlined } from '@ant-design/icons';
import {
  type ActionType,
  ModalForm, ProForm,ProFormSelect,
  ProFormText,
  ProFormTextArea, ProFormUploadButton,ProFormSwitch,
} from '@ant-design/pro-components';
import {  useRequest } from '@umijs/max';
import { Button, message } from 'antd';
import React, { FC } from 'react';

import { Item } from  "@/pages/product/item/service/data";
import {createItem} from "@/pages/product/item/service/service";

interface CreateFormProps {
  reload?: ActionType['reload'];
}

const CreateForm: FC<CreateFormProps> = (props) => {
  const { reload } = props;

  const [messageApi, contextHolder] = message.useMessage();


  const { run, loading } = useRequest(createItem, {
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
        width="600px"
        modalProps={{ okButtonProps: { loading } }}
        onFinish={async (value) => {
          value.status = value.status?1:0;
          await run({ data: value as Item });

          return true;
        }}
      >
        <ProForm.Group>
          <ProFormUploadButton
            name="avatar"
            label="图片"
            max={1}
            fieldProps={{
              name: 'file',
              listType: 'picture-card',
            }}
            action="/upload.do"
            extra="图片不能大于1M"
          />
        </ProForm.Group>

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
            placeholder="请输入标识"
            rules={[
              {
                required: true,
                message: '不能为空',
              },
            ]}
          />


          <ProFormTextArea
            width="md"
            name="desc"
            label="概略"
            placeholder="请输入"
          />


          <ProFormSelect
            initialValue="card"
            options={[
              {
                value: 'card',
                label: '卡',
              },
            ]}
            placeholder="请输入"
            width="md"
            name="type"
            label="类型"
          />



        </ProForm.Group>
        <ProForm.Group>
          <ProFormText
            width="md"
            name="duration"
            label="时长"
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
            name="length"
            label="单次时长"
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
            name="count"
            label="次数"
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
            name="price"
            label="价格"
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
            name="tagId"
            label="标签"
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
            checkedChildren="有效"
            unCheckedChildren="无效"
          />

        </ProForm.Group>
      </ModalForm>
    </>
  );
};

export default CreateForm;
