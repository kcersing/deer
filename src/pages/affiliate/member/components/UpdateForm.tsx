import {
  ProForm,
  ProFormCascader,
  ProFormDateRangePicker,
  ProFormDigit,
  ProFormList,
  ProFormMoney,
  ProFormTreeSelect,
  ProFormDateTimePicker,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
  StepsForm,
  ModalForm,
  ProFormSwitch,
  ProFormRadio,
  ProFormDatePicker,ProFormSlider,ProFormUploadButton
} from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Form, message,Slider } from 'antd';
import React, { cloneElement, useCallback, useState } from 'react';

import { Member } from  "@/pages/affiliate/member/service/data";
import {updateMember} from "@/pages/affiliate/member/service/service";

import dayjs from 'dayjs';
import type { Dayjs } from 'dayjs';

export type ModalForm = {
  trigger?: React.ReactElement<any>;
  onOk?: () => void;
  values: Partial<Member>;
};

const UpdateForm: React.FC<ModalForm> = (props) => {
  const { onOk, values, trigger } = props;
  const defaultBirthday = dayjs('1980-01-01');
  const [open, setOpen] = useState(false);

  const [messageApi, contextHolder] = message.useMessage();

  const { run } = useRequest(updateMember, {
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

        onFinish={onFinish}
      >
        <ProForm.Group>

          <ProFormUploadButton
            name="avatar"
            label="头像"
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
            name="mobile"
            label="手机号"
            placeholder="请输入手机号"
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
            width="md"
            label="性别"
            name="gender"
            initialValue={0}
            options={[
              {
                label: '男',
                value: 1,
              },
              {
                label: '女',
                value: 2,
              },
              {
                label: '未知',
                value: 0,
              },
            ]}
            rules={[
              {
                required: true,
                message: '不能为空',
              },
            ]}
          />
          <ProFormDatePicker
            width="md"
            name="birthday"
            label="出生日期"
            initialValue={defaultBirthday}
          />

          <ProFormSwitch
            name="status"
            width="md"
            label="状态"
            checkedChildren="有效" unCheckedChildren="无效"
            defaultChecked
          />
        </ProForm.Group>


        <ProForm.Group>
          <ProFormSlider
            fieldProps={{
              styles: {root:{ width: 460}},
            }}
            name="intention"
            label="意向"
            marks={{
              0: '无意向',
              20: '20%',
              40: '40%',
              60: '60%',
              80: '80%',
              100: '确定',
            }}
          />
        </ProForm.Group>
      </ModalForm>

    </>
  );
};

export default UpdateForm;
