import {
  ProForm,
  ProFormCascader,
  ProFormDatePicker,
  ProFormDateRangePicker,ProFormSwitch,
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
  ModalForm, ProFormUploadButton,
} from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Form, message } from 'antd';
import React, { cloneElement, useCallback, useState } from 'react';
import WangEditor from '@/pages/components/wangeditor'
import { updateProduct }from "@/pages/product/list/service/service";
import { Product } from  "@/pages/product/list/service/data";
import {getDicthtList} from "../../system/dict/service/service";


export type ModalForm = {
  trigger?: React.ReactElement<any>;
  onOk?: () => void;
  values: Partial<Product>;
};

const UpdateForm: React.FC<ModalForm> = (props) => {
  const { onOk, values, trigger } = props;
  const [detailBody, setDetailBody] = useState(values.content);
  const optionDetail = (data: React.SetStateAction<string>) => {
    setDetailBody(data)
  };

  const [open, setOpen] = useState(false);

  const [messageApi, contextHolder] = message.useMessage();

  const { data: tagEnum } = useRequest(() => {
    return  getDicthtList({'dictId':2})
  });
  const { run } = useRequest(updateProduct, {
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

      values.status = values.status?1:0;
      values.signSalesAt = values.salesAt[0]
      values.endSalesAt = values.salesAt[1]




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
          <ProFormUploadButton
            name="pic"
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
        </ProForm.Group>

        <ProForm.Group>
          <ProFormSelect
            width="md"
            name="tagId"
            label="标签"
            options= {tagEnum}
            fieldProps= {
              {
                showSearch: true,
                multiple: true,
                mode:"multiple",
                fieldNames:{ label: 'title', value: 'id' }
              }
            }
          />
        </ProForm.Group>

        <ProForm.Group>
          <ProFormSwitch
            name="status"
            width="md"
            label="状态"
            checkedChildren="发布"
            unCheckedChildren="下架"
            initialValue="发布"
          />

        </ProForm.Group>

        <ProForm.Group>
          <WangEditor optionDetail={optionDetail} detailBody={detailBody}/>
        </ProForm.Group>>

      </ModalForm>

    </>
  );
};

export default UpdateForm;
