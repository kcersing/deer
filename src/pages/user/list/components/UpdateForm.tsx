import {
  ProFormDateTimePicker,
  ProFormRadio,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
  StepsForm,
} from '@ant-design/pro-components';
import { FormattedMessage, useIntl, useRequest } from '@umijs/max';
import { Modal, message } from 'antd';
import React, { cloneElement, useCallback, useState } from 'react';

import {
  ProForm,
  ProFormCascader,
  ProFormDatePicker,
  ProFormDateRangePicker,
  ProFormDigit,
  ProFormList,
  ProFormMoney,

  ProFormTreeSelect,
} from '@ant-design/pro-components';
export type FormValueType = {
  target?: string;
  template?: string;
  type?: string;
  time?: string;
  frequency?: string;
} & Partial<API.RuleListItem>;

export type UpdateFormProps = {
  trigger?: React.ReactElement<any>;
  onOk?: () => void;
  values: Partial<API.RuleListItem>;
};

const UpdateForm: React.FC<UpdateFormProps> = (props) => {
  const { onOk, values, trigger } = props;

  const intl = useIntl();

  const [open, setOpen] = useState(false);

  const [messageApi, contextHolder] = message.useMessage();

  const { run } = useRequest(updateRule, {
    manual: true,
    onSuccess: () => {
      messageApi.success('Configuration is successful');
      onOk?.();
    },
    onError: () => {
      messageApi.error('Configuration failed, please try again!');
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

  return (
    <>
      {contextHolder}
      {trigger
        ? cloneElement(trigger, {
            onClick: onOpen,
          })
        : null}
      <ProForm
        stepsProps={{
          size: 'small',
        }}
        stepsFormRender={(dom, submitter) => {
          return (
            <Modal
              width={640}
              bodyStyle={{ padding: '32px 40px 48px' }}
              destroyOnClose
              title={intl.formatMessage({
                id: 'pages.searchTable.updateForm.ruleConfig',
                defaultMessage: '规则配置',
              })}
              open={open}
              footer={submitter}
              onCancel={onCancel}
            >
              {dom}
            </Modal>
          );
        }}
        onFinish={onFinish}
      >

          <ProFormText
            name="name"
            label='名称'
            width="md"
            rules={[
              {
                required: true,
                message: '请输入名称!',
              },
            ]}
          />

      </ProForm>
    </>
  );
};

export default UpdateForm;
