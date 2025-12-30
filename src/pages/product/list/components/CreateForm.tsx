import { PlusOutlined } from '@ant-design/icons';
import {
  type ActionType,
  ModalForm,
  ProForm,
  ProFormDateRangePicker,ProFormSelect,
  ProFormText,
  ProFormTextArea,
  ProFormUploadButton,ProFormList,ProCard,ProFormDependency,
  ProFormSwitch,
} from '@ant-design/pro-components';
import {  useRequest } from '@umijs/max';
import { Button, message,Descriptions, } from 'antd';
import React, { FC } from 'react';

import { createProduct }from "@/pages/product/list/service/service";
import { Product } from  "@/pages/product/list/service/data";
import dayjs from 'dayjs';
import type { Dayjs } from 'dayjs';

interface CreateFormProps {
  reload?: ActionType['reload'];
}

import Attributes from './Attributes';


const CreateForm: FC<CreateFormProps> = (props) => {
  const { reload } = props;

  const [messageApi, contextHolder] = message.useMessage();

  const defaultValue = [dayjs().startOf('month'), dayjs().startOf('month').add(2, 'month').subtract(1, 'day'),];

  const { run, loading } = useRequest(createProduct, {
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
          <Button type="primary" icon={<PlusOutlined />}>新建</Button>
        }
        width="900px"
        modalProps={{ okButtonProps: { loading } }}
        onFinish={async (value) => {
          value.status = value.status?1:0;
          await run({ data: value as Product });
          return true;
        }}
      >

        <ProForm.Group>
          <ProFormUploadButton
            name="pic"
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
        </ProForm.Group>
        <ProForm.Group>

          <ProFormText
            width="md"
            name="stock"
            label="库存"
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

          <ProFormDateRangePicker
            name="salesAt"
            label="销售时间"
            initialValue={defaultValue}
            rules={[
              {
                required: true,
                message: '不能为空',
              },
            ]}
          />
        </ProForm.Group>

      <ProForm.Group>
        <ProFormTextArea
          width="md"
          name="desc"
          label="概略"
          placeholder="请输入"
        />
      </ProForm.Group>

        <ProForm.Group>

          <Attributes />

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
