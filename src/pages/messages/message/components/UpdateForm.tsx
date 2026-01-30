import {
  ProForm,
  ProFormCascader,
  ProFormDatePicker,
  ProFormDateRangePicker,
  ProFormDigit,
  ProFormList,
  ProFormMoney,
  ProFormTreeSelect,
  ProFormDateTimePicker,
  ProFormRadio,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,ProFormSwitch,
  StepsForm,
  ModalForm,
  ProFormUploadButton,
} from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Form, message } from 'antd';
import React, { cloneElement, useCallback, useState } from 'react';

import { Item } from  "@/pages/product/item/service/data";
import {updateItem} from "@/pages/product/item/service/service";

export type ModalForm = {
  trigger?: React.ReactElement<any>;
  onOk?: () => void;
  values: Partial<Item>;
};

const UpdateForm: React.FC<ModalForm> = (props) => {
  const { onOk, values, trigger } = props;
  values.status = (values.status==1?"有效":"无效")

  console.log(values)
  const [open, setOpen] = useState(false);

  const [messageApi, contextHolder] = message.useMessage();

  const { run } = useRequest(updateItem, {
    manual: true,
    onSuccess: () => {
      messageApi.success('提交成功');
      onOk?.();
    },
    onError: () => {
      messageApi.error('提交失败，请重试！');
    },
  });

  const onCancel = useCallback(() => {
    setOpen(false);
  }, []);

  const onOpen = useCallback(() => {
    setOpen(true);
  }, []);

  const onFinish= (e) => useCallback(
    async (values?: any) => {
      values.id = e.id;
      values.status = values.status?1:0;
      await run({ data: values });
      onCancel();
    },
    [onCancel, run],
  );
  const [form] = Form.useForm<{ name: string; company: string }>();
  return (
    <>
      {contextHolder}
      {trigger
        ? cloneElement(trigger, {
            onClick: onOpen,
          })
        : null}

      <ModalForm<{
        name: string;
        company: string;
      }>
        initialValues={values}
        title="更新"
        form={form}
        autoFocusFirstInput
        modalProps={{
          destroyOnClose: true,
          onCancel: () =>{onCancel()},
        }}

        style={{ padding: '32px 40px 48px' }}
        width="600px"
        open={open}

        onFinish={onFinish(values)}
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

        </ProForm.Group>
        <ProForm.Group>
          <ProFormSelect
            initialValue="card"
            options={[
              {
                value: 'course',
                label: '课',
              },
              {
                value: 'card',
                label: '卡',
              },
            ]}
            placeholder="请输入"
            width="md"
            name="type"
            onChange={(value) => {
              setType(value);
            }}
            label="类型"
          />

        </ProForm.Group>

        <ProForm.Group>
          <ProFormDigit
            width="md"
            min={1}
            max={10}
            fieldProps={{ precision: 0 , suffix:type=="card"?"天":"节/天" }}
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
          <ProFormDigit
            min={1}
            max={10}
            fieldProps={{ precision: 0 , suffix:"分钟"}}
            width="md"
            name="length"
            hidden={ type=="card"}
            label="单次时长"
            placeholder="请输入"

          />

          <ProFormDigit
            min={1}
            max={10}
            fieldProps={{ precision: 0 , suffix:"次"}}
            width="md"
            name="count"
            label="次数"
            placeholder="请输入"

          />

          <ProFormDigit
            width="md"
            fieldProps={{ suffix:"元"}}
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

        {/*<ProForm.Group>*/}
        {/*  <ProFormText*/}
        {/*    width="md"*/}
        {/*    name="tagId"*/}
        {/*    label="标签"*/}
        {/*    placeholder="请输入"*/}
        {/*    rules={[*/}
        {/*      {*/}
        {/*        required: true,*/}
        {/*        message: '不能为空',*/}
        {/*      },*/}
        {/*    ]}*/}
        {/*  />*/}
        {/*</ProForm.Group>*/}
        <ProForm.Group>
          <ProFormTextArea
            width="md"
            name="desc"
            label="概略"
            placeholder="请输入"
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

export default UpdateForm;
