import { PlusOutlined } from '@ant-design/icons';
import {
  type ActionType,
  ModalForm, ProForm,
  ProFormText,
  ProFormTextArea,
  ProFormSwitch,
  ProFormRadio,
  ProFormDatePicker,ProFormSelect,ProFormSlider,ProFormUploadButton
} from '@ant-design/pro-components';
import {  useRequest } from '@umijs/max';
import { Button, message,Slider } from 'antd';
import React, { FC } from 'react';

import { Member } from  "@/pages/affiliate/member/service/data";
import { createMember } from "@/pages/affiliate/member/service/service";

interface CreateFormProps {
  reload?: ActionType['reload'];
}

import dayjs from 'dayjs';
import type { Dayjs } from 'dayjs';

const CreateForm: FC<CreateFormProps> = (props) => {
  const { reload } = props;

  const [messageApi, contextHolder] = message.useMessage();

  const defaultBirthday = dayjs('1980-01-01');

  const { run, loading } = useRequest(createMember, {
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
          await run({ data: value as Member });

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
        <ProForm.Group >
          <ProFormSlider

            fieldProps={{
              styles: {root:{ width: 460}},
            }}

            // size={{ xs: 24, sm: 32, md: 40, lg: 64, xl: 80, xxl: 100 }}
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

export default CreateForm;
