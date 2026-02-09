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


import { Messages } from  "./../service/data";
import {updateSend} from "./../service/service";
import WangEditor from '@/pages/components/wangeditor'

export type ModalForm = {
  trigger?: React.ReactElement<any>;
  onOk?: () => void;
  values: Partial<Messages>;
};

const UpdateForm: React.FC<ModalForm> = (props) => {
  const { onOk, values, trigger } = props;

  // const [detail, setDetail] = useState(values.content);
  const [detailBody, setDetailBody] = useState(values.content);
  const optionDetail = (data: React.SetStateAction<string>) => {

    setDetailBody(data)
  };

  const [open, setOpen] = useState(false);

  const [messageApi, contextHolder] = message.useMessage();

  const { run } = useRequest(updateSend, {
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

  const onFinish= (e) => (
    async (v?: any) => {

      var data={...e,...v}
      data.content = detailBody
      await run({ data: data });

      onCancel();
    }
  );
  const [form] = Form.useForm();
  return (
    <>
      {contextHolder}
      {trigger
        ? cloneElement(trigger, {
            onClick: onOpen,
          })
        : null}

      <ModalForm

       initialValues={{...values}}
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
            label="消息主题"
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
          <WangEditor optionDetail={optionDetail} detailBody={detailBody}/>
        </ProForm.Group>

        <ProForm.Group>
          <ProFormSelect
            name="status"
            label="状态"
            width="md"
            options={[
              { label: '草稿', value: 0 },
              { label: '已发布/发送完成', value:1 },
              { label: '定时发布中', value: 2 },
              { label: '已撤销', value: 3 },
              { label: '已归档', value: 4 },
              { label: '已删除', value: 5 },
            ]}
            placeholder="请选择"
            rules={[{ required: true, message: '请选择!' }]}
          />
        </ProForm.Group>

      </ModalForm>

    </>
  );
};

export default UpdateForm;
