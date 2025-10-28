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
  ProFormTextArea,
  StepsForm,
  ModalForm,
} from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Form, message } from 'antd';
import React, { cloneElement, useCallback, useState } from 'react';
import { updateMenu } from '@/services/ant-design-pro/menu';


export type ModalForm = {
  trigger?: React.ReactElement<any>;
  onOk?: () => void;
  values: Partial<API.Menu>;
};

const UpdateForm: React.FC<ModalForm> = (props) => {
  const { onOk, values, trigger } = props;

  const [open, setOpen] = useState(false);

  const [messageApi, contextHolder] = message.useMessage();

  const { run } = useRequest(updateMenu, {
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

  const onFinish = useCallback(
    async (values?: any) => {
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
        width="400px"
        open={open}

        onFinish={onFinish}
      >
        <ProForm.Group>
          <ProFormText
            width="md"
            name="name"
            label="名称"
            tooltip="最长为 24 位"
            placeholder="请输入名称"
          />

          <ProFormText
            width="md"
            name="title"
            label="标题"
            placeholder="请输入标题"
          />
        </ProForm.Group>
        <ProForm.Group>
          <ProFormText
            width="md"
            name="icon"
            label="图标"
            placeholder="请输入图标"
          />
          <ProFormText
            width="md"
            name="path"
            label="路由路径"
            placeholder="请输入路由路径"
          />

          <ProFormText
            width="md"
            name="component"
            label="组件路径"
            placeholder="请输入组件路径"
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

export default UpdateForm;
