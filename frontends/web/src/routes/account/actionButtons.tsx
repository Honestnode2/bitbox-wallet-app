/**
 * Copyright 2022 Shift Crypto AG
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';
import style from './account.module.css';

type TProps = {
    canSend?: boolean;
    code: string;
    exchangeBuySupported?: boolean;
}

export const ActionButtons = ({ canSend, code, exchangeBuySupported }: TProps) => {
  const { t } = useTranslation();
  return (
    <div className={style.actionsContainer}>
      {canSend ? (
        <Link key="sendLink" to={`/account/${code}/send`} className={style.send}>
          <span>{t('button.send')}</span>
        </Link>
      ) : (
        <span key="sendDisabled" className={`${style.send} ${style.disabled}`}>
          {t('button.send')}
        </span>
      )}
      <Link key="receive" to={`/account/${code}/receive`} className={style.receive}>
        <span>{t('button.receive')}</span>
      </Link>
      { exchangeBuySupported && (
        <Link key="buy" to={`/buy/info/${code}`} className={style.buy}>
          <span>{t('button.buy')}</span>
        </Link>
      )}
    </div>
  );
};
