import React from "react";
import { NextComponentType, NextPageContext } from "next";
import Image from 'next/image';

import Grid from "@mui/material/Grid";
import { styled } from '@mui/material/styles';
import LinearProgress, { linearProgressClasses } from '@mui/material/LinearProgress';

const BorderLinearProgress = styled(LinearProgress)(({ theme }) => ({
  height: 25,
  borderRadius: 12,
  [`&.${linearProgressClasses.colorPrimary}`]: {
    backgroundColor: theme.palette.grey[theme.palette.mode === 'light' ? 200 : 800],
  },
  [`& .${linearProgressClasses.bar}`]: {
    borderRadius: 12,
    backgroundColor: '#f39800',
  },
  ":before": {
    content: '"Now loading..."',
    position: "absolute",
    "z-index": 100,
    inset: 0,
    margin: "auto",
    fontFamily: "'Zen Maru Gothic', sans-serif",
    fontWeight: "500",
  }
}));

const _Loading: NextComponentType<
    NextPageContext,
    Record<string, unknown>
> = () => {
    return (
      <Grid container alignItems="center" justifyContent="center" height="100dvh" className="fadeIn" style={{ width: '100%' }}>
          <Grid item style={{ width: '100%' }}>
              <Grid container justifyContent="center" alignItems="center" direction="column" spacing="48">
                  <Grid item style={{ maxWidth: '400px', width: '80%', maxHeight: '400px', height: '60dvw' }}>
                    <div style={{ position: 'relative', width: '100%', height: '100%', textAlign: 'center' }}>
                        <Image src="/logo.png" alt="logo" layout="fill" objectFit="contain" />
                    </div>
                  </Grid>
                  <Grid item textAlign="center" style={{ width: '80%' }}>
                      <BorderLinearProgress />
                  </Grid>
              </Grid>
          </Grid>
      </Grid>
    );
};

const Loading = React.memo(_Loading);
export default Loading;
