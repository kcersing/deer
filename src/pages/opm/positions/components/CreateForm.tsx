import { PlusOutlined } from '@ant-design/icons';
import {
  type ActionType,
  ModalForm, ProForm,ProFormSwitch,
  ProFormText,
  ProFormTextArea,ProFormSelect
} from '@ant-design/pro-components';
import {  useRequest } from '@umijs/max';
import { Button, message } from 'antd';
import React, { FC } from 'react';

import { Positions } from  "@/pages/opm/positions/service/data";
import {createPositions} from "@/pages/opm/positions/service/service";
import {getDepartmentsList} from "../../departments/service/service";
import {getPositionsList} from "../service/service";

interface CreateFormProps {
  reload?: ActionType['reload'];
}

const CreateForm: FC<CreateFormProps> = (props) => {
  const { reload } = props;

  const [messageApi, contextHolder] = message.useMessage();


  const { run, loading } = useRequest(createPositions, {
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
       title='新建菜单'
        trigger={
          <Button type="primary" icon={<PlusOutlined />}>
            新建
          </Button>
        }
        width="800px"
        modalProps={{ okButtonProps: { loading } }}
        onFinish={async (value) => {
          value.status = value.status?1:0;
          await run({ data: value as Positions });

          return true;
        }}
      >


        <ProForm.Group>
          <ProFormSelect
            name="parentId"
            label="上级"
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
          <ProFormText
            width="md"
            name="name"
            label="职位名称"
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
          <ProFormSelect
            name="departmentId"
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
        </ProForm.Group>


        <ProForm.Group>
          <ProFormTextArea
            width="md"
            name="desc"
            label="概略"
            placeholder="请输入"
          />
          <ProFormText
            width="md"
            name="quota"
            label="编制人数"
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

export default CreateForm;
