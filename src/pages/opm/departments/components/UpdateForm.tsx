import {
  ProForm,
  ProFormText,ProFormSwitch,
  ModalForm,ProFormSelect,ProFormTextArea
} from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Form, message } from 'antd';
import React, { cloneElement, useCallback, useState } from 'react';

import { Departments } from  "@/pages/opm/departments/service/data";
import {updateDepartments} from "@/pages/opm/departments/service/service";
import {getDepartmentsList} from "../service/service";

export type ModalForm = {
  trigger?: React.ReactElement<any>;
  onOk?: () => void;
  values: Partial<Departments>;
};

const UpdateForm: React.FC<ModalForm> = (props) => {
  const { onOk, values, trigger } = props;

  const [open, setOpen] = useState(false);

  const [messageApi, contextHolder] = message.useMessage();

  const { run } = useRequest(updateDepartments, {
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
          <ProFormSelect
            name="parentId"
            label="上级"
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
        </ProForm.Group>
        <ProForm.Group>
          <ProFormText
            width="md"
            name="name"
            label="部门名称"
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
          <ProForm.Group>
            <ProFormSelect
              name="managerId"
              label="负责人ID"
              width="sm"
              params={{current: 999, pageSize: 1}}
              request={(params)=>{
                return getUserList({params}).then((res) => {return res.data})
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
          <ProFormSwitch
            name="status"
            width="md"
            label="状态"
            checkedChildren="有效"
            unCheckedChildren="无效"
            initialValue="有效"
          />

        </ProForm.Group>
      </ModalForm>

    </>
  );
};

export default UpdateForm;
