import { PlusOutlined } from '@ant-design/icons';
import {
  type ActionType,
  ModalForm, ProForm,
  ProFormText,
  ProFormSwitch,
  ProFormDatePicker,
  ProFormUploadButton,ProFormRadio
} from '@ant-design/pro-components';
import {  useRequest } from '@umijs/max';
import { Button, message } from 'antd';
import React, { FC,useState } from 'react';

import { User } from  "@/pages/opm/user/service/data";
import {createUser} from "@/pages/opm/user/service/service";

import WangEditor from '@/pages/components/wangeditor'

interface CreateFormProps {
  reload?: ActionType['reload'];
}

const CreateForm: FC<CreateFormProps> = (props) => {
  const { reload } = props;

  const [messageApi, contextHolder] = message.useMessage();

  const [detail, setDetail] = useState('');
  const [detailBody, setDetailBody] = useState('');


  const optionDetail = (data: React.SetStateAction<string>) => {
    setDetail(data)
  };


  const { run, loading } = useRequest(createUser, {
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
          value.status = value.status?1:0;
          value.detail=detail
          await run({ data: value as User });
          return true;
        }}
      >


        <ProForm.Group>
          <ProFormUploadButton
            max={1}
            width="sm"
            name="avatar"
            label="头像"
            fieldProps={{
              name: 'file',
              listType: 'picture-card',
            }}
            action="/upload.do"
            extra="头像备注"
          />
        </ProForm.Group>
        {/*number | "xs" | "sm" | "md" | "lg" | "xl"*/}
        <ProForm.Group>
          <ProFormText
            width="sm"
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
            width="sm"
            name="username"
            label="登录账号"
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
            width="sm"
            name="mobile"
            label="手机号"
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

          <ProFormRadio.Group
            width="sm"
            name="gender"
            label="性别"
            options={[
              { value: 1, label: '男' },
              { value: 2, label: '女' },
              { value: 3, label: '保密 ' },
            ]}
          />

          <ProFormDatePicker
            width="sm"
            name="birthday"
            label="出生日期"
          />
          <ProFormSwitch
            width="sm"
            name="status"
            label="状态"
            checkedChildren="开启"
            unCheckedChildren="关闭"
          />

        </ProForm.Group>
        <ProForm.Group>
          <WangEditor optionDetail={optionDetail} detailBody={detailBody}/>
        </ProForm.Group>
      </ModalForm>
    </>
  );
};
export default CreateForm;
