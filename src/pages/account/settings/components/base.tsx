import { UploadOutlined } from '@ant-design/icons';
import {
  ProForm,
  ProFormDependency,
  ProFormFieldSet,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
  ProFormRadio,ProFormDatePicker,
} from '@ant-design/pro-components';
import { useRequest } from '@umijs/max';
import { Button, Input, message, Upload } from 'antd';
import React from 'react';
import { queryCity, queryCurrent, queryProvince } from '../service';
import useStyles from './index.style';
import { User } from  "@/pages/account/service/data";
import {getUser,updateUser} from "@/pages/account/service/service";

import dayjs from 'dayjs';
import type { Dayjs } from 'dayjs';

const BaseView: React.FC = () => {


  const { styles } = useStyles();
  // 头像组件 方便以后独立，增加裁剪之类的功能
  const AvatarView = ({ avatar }: { avatar: string }) => (
    <>
      <div className={styles.avatar_title}>头像</div>
      <div className={styles.avatar}>
        <img src={avatar} alt="avatar" />
      </div>
      <Upload showUploadList={false}>
        <div className={styles.button_view}>
          <Button>
            <UploadOutlined />
            更换头像
          </Button>
        </div>
      </Upload>
    </>
  );
  const { data: currentUser, loading } = useRequest(() => {
    return getUser({id:0});
  });
  const getAvatarURL = () => {
    if (currentUser) {
      if (currentUser.avatar) {
        return currentUser.avatar;
      }
      const url =
        'https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png';
      return url;
    }
    return '';
  };

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


  const handleFinish = async (values?: any) => {



    values.province = values.province.value

    var  data={...currentUser,...values}
console.log(data)
    // await run({ data: data });
  };


  return (
    <div className={styles.baseView}>
      {loading ? null : (
        <>
          <div className={styles.left}>
            <ProForm
              layout="vertical"
              onFinish={handleFinish}
              submitter={{
                searchConfig: {
                  submitText: '更新基本信息',
                },
                render: (_, dom) => dom[1],
              }}
              initialValues={{
                ...currentUser,
              }}
              hideRequiredMark
            >

              <ProFormText
                width="md"
                name="name"
                label="姓名"

                rules={[
                  {
                    required: true,
                    message: '请输入您的姓名!',
                  },
                ]}
              />

              <ProFormRadio.Group
                width="md"
                name="gender"
                label="性别"
                rules={[
                  {
                    required: true,
                    message: '请输入您的性别!',
                  },
                ]}
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
                    value: 3,
                  },
                ]}
              />

              <ProFormDatePicker
                width="md"
                name="birthday"
                label="出生日期"
              />

              <ProFormTextArea
                name="desc"
                label="个人简介"
                rules={[
                  {
                    required: true,
                    message: '请输入个人简介!',
                  },
                ]}
                placeholder="个人简介"
              />
              <ProFormText
                width="md"
                name="email"
                label="邮箱"
              />

              <ProForm.Group title="所在省市" size={8}>
                <ProFormSelect
                  width="sm"
                  fieldProps={{
                    labelInValue: true,
                  }}
                  name="province"
                  request={async () => {
                    return queryProvince().then(({ data }) => {
                      return data.map((item) => {
                        return {
                          label: item.name,
                          value: item.id,
                        };
                      });
                    });
                  }}
                />
                <ProFormDependency name={['province']}>
                  {({ province }) => {
                    return (
                      <ProFormSelect
                        params={{
                          key: province?.value,
                        }}
                        name="city"
                        width="sm"
                        disabled={!province}
                        request={async () => {
                          if (!province?.key) {
                            return [];
                          }
                          return queryCity(province.key || '').then(
                            ({ data }) => {
                              return data.map((item) => {
                                return {
                                  label: item.name,
                                  value: item.id,
                                };
                              });
                            },
                          );
                        }}
                      />
                    );
                  }}
                </ProFormDependency>
              </ProForm.Group>
              <ProFormText
                width="md"
                name="address"
                label="街道地址"
              />
            </ProForm>
          </div>
          <div className={styles.right}>
            <AvatarView avatar={getAvatarURL()} />
          </div>
        </>
      )}
    </div>
  );


};
export default BaseView;
