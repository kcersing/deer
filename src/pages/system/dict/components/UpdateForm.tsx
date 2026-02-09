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
  ProFormSwitch,
  StepsForm,
  ModalForm,
} from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Form, message } from 'antd';
import React, { cloneElement, useCallback, useState } from 'react';

import { Dictht,Dict } from  "@/pages/system/dict/service/data";
import {updateDict}  from "@/pages/system/dict/service/service";

export type ModalForm = {
  trigger?: React.ReactElement<any>;
  onOk?: () => void;
  values: Partial<Dictht>;
};

const UpdateForm: React.FC<ModalForm> = (props) => {
  const { onOk, values, trigger } = props;

  const [open, setOpen] = useState(false);

  const [messageApi, contextHolder] = message.useMessage();

  const { run } = useRequest(updateDict, {
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
        width="800px"
        open={open}

        onFinish={onFinish(values)}
      >
        <ProForm.Group>
          <ProFormText
            width="md"
            name="title"
            label="名称"
            tooltip="最长为 24 位"
            placeholder="请输入名称"
          />

          <ProFormText
            width="md"
            name="code"
            label="标识"
            disabled
            placeholder="请输入标识"
          />

        </ProForm.Group>
        <ProForm.Group>

          <ProFormTextArea
            width="md"
            name="desc"
            label="概略"
            placeholder="请输入"
          />

          <ProFormSwitch
            name="status"
            width="md"
            label="状态"
            checkedChildren="开启"
            unCheckedChildren="关闭"
          />

        </ProForm.Group>
      </ModalForm>

    </>
  );
};

export default UpdateForm;
