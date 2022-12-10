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
      <Grid container alignItems="center" justifyContent="center" height="100dvh" className="fadeIn">
          <Grid item>
              <Grid container justifyContent="center" direction="column" spacing="48">
                  <Grid item textAlign="center">
                    <Image src="/logo.png" alt="logo" width={500} height={300} />
                  </Grid>
                  <Grid item textAlign="center">
                      <BorderLinearProgress />
                  </Grid>
              </Grid>
          </Grid>
      </Grid>
    );
};

const Loading = React.memo(_Loading);
export default Loading;
