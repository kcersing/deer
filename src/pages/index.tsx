import { history } from '@umijs/max';
import { Button, Card, Result } from 'antd';
import React from 'react';
import { useLocation } from 'umi';
export default function Page() {
  let location = useLocation();
  return (
    <div>
      { history.push('/welcome') }
    </div>
  );
}
