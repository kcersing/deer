import {
  ProForm,
  ProFormText,ProFormSwitch,
  ModalForm,ProFormUploadButton,ProFormRadio,ProFormDatePicker,ProFormSelect
} from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Form, message } from 'antd';
import React, { cloneElement, useCallback, useState } from 'react';
import WangEditor from '@/pages/components/wangeditor'
import { User } from  "@/pages/opm/user/service/data";
import {updateUser} from "@/pages/opm/user/service/service";
import {getMessagesTypes} from "../../../messages/message/service/service";
import {getPositionsList} from "../../positions/service/service";
import {getDepartmentsList} from "../../departments/service/service";

export type ModalForm = {
  trigger?: React.ReactElement<any>;
  onOk?: () => void;
  values: Partial<User>;
};

const UpdateForm: React.FC<ModalForm> = (props) => {
  const { onOk, values, trigger } = props;
  const [detailBody, setDetailBody] = useState(values.content);
  const optionDetail = (data: React.SetStateAction<string>) => {
    setDetailBody(data)
  };
  const [open, setOpen] = useState(false);

  const [messageApi, contextHolder] = message.useMessage();

  const { run } = useRequest(updateUser, {
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

        {/*<ProForm.Group>*/}
        {/*  <ProFormUploadButton*/}
        {/*    max={1}*/}
        {/*    width="sm"*/}
        {/*    name="avatar"*/}
        {/*    label="头像"*/}
        {/*    fieldProps={{*/}
        {/*      name: 'file',*/}
        {/*      listType: 'picture-card',*/}
        {/*    }}*/}
        {/*    action="/upload.do"*/}
        {/*    extra="头像备注"*/}
        {/*  />*/}
        {/*</ProForm.Group>*/}

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

        </ProForm.Group>

        <ProForm.Group>
          <ProFormText
            width="sm"
            name="username"
            label="登录账号"
            placeholder="请输入"
            disabled
          />
          <ProFormText
            width="sm"
            name="mobile"
            label="手机号"
            disabled
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

        </ProForm.Group>

        <ProForm.Group>
          <ProFormSwitch
            width="sm"
            name="status"
            label="状态"
            checkedChildren="开启"
            unCheckedChildren="关闭"
            initialValue="开启"
          />
        </ProForm.Group>

        <ProForm.Group>
          <ProFormSelect
            name="departmentsId"
            label="部门"
            width="sm"
            params={{current: 999, pageSize: 1}}
            request={(params)=>{
              return getDepartmentsList({params}).then((res) => {return res.data})
            }}

            fieldProps={{
              fieldNames: {
                label: 'name',
                value: 'id',
              },
            }}
            placeholder="请选择"
            rules={[{ required: true, message: '请选择!' }]}
          />

          <ProFormSelect
            name="positionsId"
            label="职位"
            width="sm"
            params={{current: 999, pageSize: 1}}
            request={(params)=>{
              return getPositionsList({params}).then((res) => {return res.data})
            }}

            fieldProps={{
              fieldNames: {
                label: 'name',
                value: 'id',
              },
            }}
            placeholder="请选择"
            rules={[{ required: true, message: '请选择!' }]}
          />
        </ProForm.Group>


        <ProForm.Group>
          <WangEditor optionDetail={optionDetail} detailBody={detailBody}/>
        </ProForm.Group>

      </ModalForm>

    </>
  );
};

export default UpdateForm;
